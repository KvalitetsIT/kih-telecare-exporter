package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend/kih/exporttypes"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend/oioxds"
	"github.com/KvalitetsIT/kih-telecare-exporter/internal"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	backendImpl   string
	patient       string
	file          string
	source        string
	kihcreatedby  string
	kihsosiserver string
	usesosi       bool
	setnow        bool
	date          string
	kihurl        string
	xdsgenerator  string
	xdsserver     string
)

func init() {
	rootCmd.AddCommand(testInjectCmd)
	// Default for when reports is to be started from
	viper.SetDefault("clinician.batchsize", 100)
	testInjectCmd.Flags().StringVarP(&backendImpl, "backend", "b", "kih", "-b indicates with exporter backend to use. Supported backends: kih,oioxds")
	testInjectCmd.Flags().StringVarP(&patient, "patient", "p", "", "-p is a path to JSON file with patient information")
	testInjectCmd.Flags().StringVarP(&file, "file", "f", "", "-f is a path to JSON file measurent data to be sent")
	testInjectCmd.Flags().StringVarP(&source, "source", "s", "", "-s is a path to directory with JSON files with measurent data to be sent")

	// KIH flags
	testInjectCmd.Flags().StringVarP(&kihcreatedby, "kihcreatedby", "", "", "Sets created by in OIO request")
	testInjectCmd.Flags().StringVarP(&kihsosiserver, "kihsosiserver", "", "", "Sets URL for SOSI Server")
	testInjectCmd.Flags().BoolVarP(&usesosi, "usesosi", "", false, "Use SOSI?")
	testInjectCmd.Flags().BoolVarP(&setnow, "setnow", "", false, "Set timestamp on measurement to now?")
	testInjectCmd.Flags().StringVarP(&date, "date", "", "", "Specify date to use? - format YYYY-mm-ddTHH:MM:ss")
	testInjectCmd.Flags().StringVarP(&kihurl, "kihurl", "", "https://kihdb-devel.oth.io/services/monitoringDataset", "Sets URL for KIHDB endpoint (https://kihdb-devel.oth.io/services/monitoringDataset)")

	// XDS Flags
	testInjectCmd.Flags().StringVarP(&xdsgenerator, "xdsgen", "", "http://localhost:9010/api/createphmr", "URL for xds generator")
	testInjectCmd.Flags().StringVarP(&xdsserver, "xdsrepo", "", "", "URL for xds Server")

	if err := testInjectCmd.MarkFlagRequired("patient"); err != nil {
		logrus.Fatalf("error setting up flags %v", err)
	}
	if err := testInjectCmd.MarkFlagRequired("kihcreatedby"); err != nil {
		logrus.Fatalf("error setting up flags %v", err)
	}
}

func setupExporterBackend(application *app.Config, dummyApi measurement.MeasurementApi) backend.ExportBackend {
	var exporter backend.ExportBackend

	switch backendImpl {
	case "oioxds":
		log.Warnf("Using OIO XDS Backend")
		log.Debugf("Use SOSI? %v", usesosi)
		application.Export.OIOXDSExport.XdsGenerator.URL = xdsgenerator

		exporter = oioxds.InitExporter(application, dummyApi)
	default:
		log.Warnf("Unsupported backend %s", backendImpl)
		os.Exit(1)
	}

	return exporter
}

// Export measurements
func exporterMeasurements(application *app.Config, expr backend.ExportBackend, measurements []measurement.Measurement) error {
	now := time.Now()
	types := expr.GetExportTypes()

	for _, v := range measurements {
		id := uuid.New()
		log.Debugf("Type - '%s'", v.Type)

		if !expr.ShouldExport(v) {
			log.Infof("Measurement is marked as not exportable - %s", v.Type)
			continue
		}

		if len(v.Type) == 0 || len(v.Measurement.Unit) == 0 {
			log.Warnf("Invalid measurement type %s ", v.Type)
			continue
		}
		if setnow {
			log.Debugf("Setting time to now - %s", now.Format(time.RFC3339))
			v.Timestamp = now
		}
		if len(date) > 0 {
			location, err := time.LoadLocation("Europe/Copenhagen")
			if err != nil {
				panic(err)
			}
			dateStamp, err := time.ParseInLocation("2006-01-02T15:04:05", date, location)
			if err != nil {
				log.Fatalf("Error parsing date %s, got %v", date, err)
			}
			log.Debugf("Using date %v", dateStamp)
			v.Timestamp = dateStamp
		}

		fmt.Printf("ID..........: %s\n", id)
		fmt.Printf("Type........: %s\n", v.Type)
		if "blood_pressure" == v.Type {
			exportType := types[v.Type]
			fmt.Printf("Actual Value: %f/%f\n", v.Measurement.Systolic, v.Measurement.Diastolic)
			t := exportType.(exporttypes.BloodPressureType)

			fmt.Printf("Sent Value..: %v/%v\n", t.GetSystolic().GetResultText(v), t.GetDiastolic().GetResultText(v))
		} else {
			fmt.Printf("Actual Value: %f\n", v.Measurement.Value)
			fmt.Printf("Sent Value..: %v\n", types[v.Type].GetResultText(v))
		}
		fmt.Printf("Unit........: %v\n", v.Measurement.Unit)
		fmt.Printf("Sent Unit...: %v\n", types[v.Type].GetResultUnitText())
		fmt.Printf("Date........: %s\n", v.Timestamp)

		res, err := expr.ConvertMeasurement(v, repository.MeasurementExportState{ID: id})
		if err != nil {
			log.Errorf("Error running exports %v", err)
			return errors.Wrap(err, "Error converting measurement")
		}

		//log.Infof("Sending - %v", res)
		serverOk, err := expr.ExportMeasurement(res)

		if err != nil {
			log.Errorf("Error exporting measurement - %v", err)
			//log.Errorf("Sent \n%s", res)
		}
		log.Debugf("Measurent exported - server said : %s", serverOk)

	}

	return nil
}

// Load measurements from file or directory
func loadMeasurements() []measurement.Measurement {
	var mes []measurement.Measurement

	// Handle file
	if len(file) > 0 {
		measure, err := loadMeasurementFromFile(file)
		if err != nil {
			log.Fatalf("Error fetching patient - %v", err)
		}
		log.Debugf("Loaded %v", measure)
		mes = append(mes, measure)

	}
	// Handle directory
	if len(source) > 0 {
		log.Warnf("Loading measurement from directory - %v", source)
		files, err := ioutil.ReadDir(source)
		if err != nil {
			log.Fatalf("Error reading source directory %s - %v", source, err)
		}
		for _, v := range files {
			if !v.IsDir() {
				if strings.HasSuffix(v.Name(), ".json") {
					log.Infof("Reading file - %s/%s", source, v.Name())
					measure, err := loadMeasurementFromFile(fmt.Sprintf("%s/%s", source, v.Name()))
					if err != nil {
						log.Fatalf("Error fetching patient - %v", err)
					}
					log.Debugf("Loaded %v", measure)
					mes = append(mes, measure)
				}
			}
		}
	}

	return mes
}

// Setup the inject commmand
var testInjectCmd = &cobra.Command{
	Use:     "testinject",
	Short:   "Reads measurements and patients from file and exports based on config",
	Aliases: []string{"ti"},
	Run: func(cmd *cobra.Command, args []string) {

		application, err := app.InitConfig()
		if err != nil {
			logrus.Fatal("Error initializing exporter ", err)
		}
		pkg := app.GetPackage(reflect.TypeOf(empty{}).PkgPath())
		log = app.NewLogger(application.GetLoggerLevel(pkg))

		dummyApi := internal.TestInjectorApi{}

		log.Debug("Patient: ", patient)

		patientInformation, err := loadPatientData(patient)
		if err != nil {
			log.Fatalf("Error fetching patient - %v", err)
		}
		log.Debugf("Using patient - %v", patientInformation)

		dummyApi.Patient = patientInformation

		exporter := setupExporterBackend(application, dummyApi)

		mes := loadMeasurements()

		if err := exporterMeasurements(application, exporter, mes); err != nil {
			log.Errorf("Error exporting measurements %v", err)
		}

	},
}

// Load Patient Data from file
func loadPatientData(p string) (measurement.PatientResult, error) {
	var pat measurement.PatientResult
	data, err := ioutil.ReadFile(p)

	if err != nil {
		return pat, errors.Wrap(err, "Error reading patient data")
	}

	if err := json.Unmarshal(data, &pat); err != nil {
		return pat, errors.Wrap(err, "Error parsing patient data")
	}
	return pat, nil
}

// Load measurement from local file
func loadMeasurementFromFile(p string) (measurement.Measurement, error) {
	var pat measurement.Measurement
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return pat, errors.Wrap(err, "Error reading patient data")
	}

	if err := json.Unmarshal(data, &pat); err != nil {
		return pat, errors.Wrap(err, "Error parsing patient data")
	}
	return pat, nil
}
