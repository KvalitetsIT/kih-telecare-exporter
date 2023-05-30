package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var config *app.Config
var log *logrus.Logger

type repositoryImpl struct {
	Value    string
	conn     *sqlx.DB
	location *time.Location
}

func InitRepository(cfg *app.Config, db *sqlx.DB) (Repository, error) {
	config = cfg
	pkg := app.GetPackage(reflect.TypeOf(repositoryImpl{}).PkgPath())
	log = app.NewLogger(config.GetLoggerLevel(pkg))
	repo := repositoryImpl{}

	repo.conn = db

	// Set the time zone

	location, err := time.LoadLocation("Europe/Copenhagen")
	if err != nil {
		log.Error("Error settting location", err)
	}

	repo.location = location
	log.Debug("Initialized repository")
	return repo, nil
}

func (mi repositoryImpl) getSession() (*sqlx.DB, error) {
	for i := 0; i < 20; i++ { // nolint
		if err := mi.conn.Ping(); err != nil {
			log.Errorf("Error pinging database - %v", err)
			time.Sleep(time.Millisecond * 100)
			i++ // nolint
		}
		return mi.conn, nil // nolint
	}

	return nil, fmt.Errorf("Error gettting session towards db")
}

func getSumFromDb(sess *sqlx.DB, query string) int {
	start := time.Now()
	var sum int

	if err := sess.Get(&sum, query); err != nil {
		log.Error("Error quering db ", err)
		return 0
	}

	log.Infof("func=getsumfromdb tt=%s", time.Since(start))
	return sum
}

func (mi repositoryImpl) GetTotals() (int, int, int, int) {
	start := time.Now()
	sess, err := mi.getSession()
	if err != nil {
		log.Errorf("Error gettting DB session - %v", err)
		return 0, 0, 0, 0
	}
	afterdb := time.Now()
	log.Infof("Spend %s on getting db connection", afterdb.Sub(start))
	total := getSumFromDb(sess, "SELECT count(measurement) from measurements")
	failed := getSumFromDb(sess, fmt.Sprintf("SELECT count(measurement) from measurements where status=%d", FAILED))
	tempfailed := getSumFromDb(sess, fmt.Sprintf("SELECT count(measurement) from measurements where status=%d", TEMP_FAILURE))
	rejected := getSumFromDb(sess, fmt.Sprintf("SELECT count(measurement) from measurements where status=%d", NO_EXPORT))

	log.Infof("func=gettotals tt=%s dbt=%s totalst=%s", time.Since(start), afterdb.Sub(start), time.Since(afterdb))
	return total, failed, tempfailed, rejected
}

// Returns time for last run, total runs,failed, successfull, status of last run
func (mi repositoryImpl) GetRuns() (time.Time, int, int, int, int) {
	start := time.Now()
	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		return time.Now(), 0, 0, 0, 0
	}
	afterdb := time.Now()
	total := getSumFromDb(sess, "SELECT count(lastrun) from runstatus")
	failed := getSumFromDb(sess, fmt.Sprintf("SELECT count(lastrun) from runstatus where status=%d", FAILED))
	successfull := getSumFromDb(sess, fmt.Sprintf("SELECT count(lastrun) from runstatus where status=%d", COMPLETED))

	var result RunStatus
	if err = sess.Get(&result, "SELECT * from runstatus ORDER BY lastrun DESC LIMIT 1"); err != nil {
		if err != sql.ErrNoRows {
			log.Error("Error getting status:", err)
		}
	}

	log.Infof("func=getruns tt=%s dbt=%s totalst=%s", time.Since(start), afterdb.Sub(start), time.Since(afterdb))
	return result.Lastrun, result.Status, total, successfull, failed
}

func (mi repositoryImpl) CheckRepository() error {
	log.Debug("Testing repository connection")
	conn, err := mi.getSession()
	if err != nil {
		return errors.Wrap(err, "Error getting DB session")
	}

	return conn.Ping()
}

// FindOrCreateMeasuremnet is used when the Initial loads is run
func (mi repositoryImpl) FindOrCreateMeasurement(m MeasurementExportState) (MeasurementExportState, error) {
	var res MeasurementExportState
	// Have we seen it before ...
	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		return MeasurementExportState{}, errors.Wrap(err, "Error getting conection")
	}
	err = sess.Get(&res, "SELECT id,measurement,patient,status,created_at,updated_at FROM measurements WHERE measurement=?", m.Measurement)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error("Error querying database", err)
			log.Infof("Trace %+v", err)

			return MeasurementExportState{}, errors.Wrap(err, "Error querying database")
		}
	}

	if len(res.Measurement) > 0 {
		return res, nil
	} else {
		newUUID := uuid.New()
		m.ID = newUUID
		//m.Status = INITIAL
		now := time.Now()
		if m.CreatedAt.Time.IsZero() {
			m.CreatedAt.Time = now
		}
		m.UpdatedAt.Time = now
		tx := sess.MustBegin()
		_, err := tx.Exec("INSERT INTO measurements (id,measurement,patient,status,created_at,updated_at) VALUES (?,?,?,?,?,?)", m.ID, m.Measurement, m.Patient, m.Status, m.CreatedAt.Time, m.UpdatedAt.Time)

		if err != nil {
			log.Error("Error inserting data ", err)
			log.Infof("Trace %+v", err)
			return m, errors.Wrap(err, "Error inserting data")
		}
		if err := tx.Commit(); err != nil {
			log.Infof("Trace %+v", err)
			return m, errors.Wrap(err, "Error commit transaction")
		}
	}

	return m, nil
}

func (mi repositoryImpl) UpdateMeasurement(m MeasurementExportState) (MeasurementExportState, error) {
	// Check that the measurement is already populated
	if len(m.ID.String()) == 0 || len(m.Measurement) == 0 {
		return m, fmt.Errorf("Measurement is not set - unknown measurement")
	}

	m.UpdatedAt.Time = time.Now()

	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		log.Infof("Trace %+v", err)
		return m, errors.Wrap(err, "Error getting conection")
	}

	tx, err := sess.Begin()
	if err != nil {
		log.Error("Error creating transaction", err)
		log.Infof("Trace %+v", err)
		return m, errors.Wrap(err, "Error creating transaction")
	}

	res, err := tx.Exec("UPDATE measurements set updated_at=?, status=? where measurement=?", m.UpdatedAt.Time, m.Status, m.Measurement)
	if err != nil {
		log.Error("Error updating row", err)
		log.Infof("Trace %+v", err)
		if rerr := tx.Rollback(); rerr != nil {
			log.Infof("Trace %+v", err)
			log.Errorf("Error rollback transaction - %v", rerr)
		}
		return m, errors.Wrap(err, "Error updating measurement")
	}

	numRows, _ := res.RowsAffected()

	log.Debug("Updates ", numRows, " rows")
	err = tx.Commit()
	if err != nil {
		return m, errors.Wrap(err, "Error commmiting tranaction")
	}

	log.Debug("M", m, " ", m.UpdatedAt.Time)

	return m, nil

}
func (mi repositoryImpl) Close() error {
	log.Info("Closing DB adapter")
	if err := mi.conn.Close(); err != nil {
		log.Error("Error closing adapter - ", err)
	}
	return nil
}

func (mi repositoryImpl) FindMeasurements() ([]MeasurementExportState, error) {
	var measurements []MeasurementExportState

	return measurements, nil
}

func (mi repositoryImpl) FindMeasurement(id string) (MeasurementExportState, error) {
	var res MeasurementExportState
	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		return MeasurementExportState{}, errors.Wrap(err, "Error getting conection")
	}
	err = sess.Get(&res, "SELECT id,measurement,patient,status,created_at,updated_at FROM measurements WHERE id=?", id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error("Error querying database", err)
			log.Infof("Trace %+v", err)

			return MeasurementExportState{}, errors.Wrap(err, "Error querying database")
		} else {
			return MeasurementExportState{}, fmt.Errorf("Id %s not found : %w", id, err)

		}
	}

	return res, nil
}

func (mi repositoryImpl) FindMeasurementsByStatus(status int) ([]MeasurementExportState, error) {
	var measurements []MeasurementExportState

	sess, err := mi.getSession()
	if err != nil {
		return measurements, errors.Wrap(err, "Error getting session")
	}

	if err := sess.Select(&measurements, "SELECT id,measurement,patient,status,created_at,updated_at FROM measurements where status=?", status); err != nil {
		return measurements, errors.Wrap(err, "Error retrieving measuremnts")
	}

	return measurements, nil
}
