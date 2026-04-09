package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func (mi repositoryImpl) StartExport() (RunStatus, error) {
	var lastRun RunStatus
	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		return lastRun, errors.Wrap(err, "Error getting conection")
	}

	var runs []int
	if err := sess.Select(&runs, "SELECT count(*) from runstatus where status=?", COMPLETED); err != nil {
		fmt.Println("HERE", err)
		return lastRun, errors.Wrap(err, "Error getting row count")
	}

	lr := RunStatus{Id: uuid.New(), Status: INITIAL}
	lr.CreatedAt.Time = time.Now()

	if len(runs) == 1 && runs[0] == 0 { // no last run found - start from the beginning
		startDate := viper.GetString("export.start")
		t, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Fatal("Error parsing start date - ", err, " ", t)
		}
		log.Warn("Using startme - ", t)
		lr.Lastrun = t

		tx, err := sess.Begin()
		if err != nil {
			return lastRun, errors.Wrap(err, "Error creating transaction")
		}
		_, err = tx.Exec("INSERT INTO runstatus (id, lastrun, status, created_at,updated_at) VALUES (?,?,?,?,?)",
			uuid.New(), t, COMPLETED, t, t)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Errorf("Error rolling back transaction %+v", rerr)
				log.Infof("Trace %+v", err)
			}
			return lastRun, errors.Wrap(err, "Error inserting row")
		}
		if err := tx.Commit(); err != nil { // commmit the transaction
			return lastRun, errors.Wrap(err, "Error commiting transaction")
		}
	} else { // Get the last one

		if err := sess.Get(&lastRun, "select * from runstatus where status=? order by lastrun desc limit 1", COMPLETED); err != nil {
			if err != sql.ErrNoRows {
				log.Errorf("Error reading from database %+v", err)
				log.Infof("Trace %+v", err)
				return lastRun, errors.Wrap(err, "Error getting lastrun")
			} else {
				log.Debug("No results found")
			}
		}

		log.Debug("Using last run for start: ", lastRun)
		delta := time.Since(lastRun.Lastrun)
		log.Debug("Minutes since last run ", delta.Minutes())
		log.Debug("Setting startime ", lastRun.Lastrun.Add(-5*time.Minute))
		lr.CreatedAt.Time = time.Now()
		lr.Lastrun = lastRun.Lastrun.Add(-30 * time.Minute)
	}

	tx, err := sess.Begin()
	if err != nil {
		return lastRun, errors.Wrap(err, "Error creating transaction")
	}
	_, err = tx.Exec("INSERT INTO runstatus (id, lastrun, status, created_at,updated_at) VALUES (?,?,?,?,?)",
		lr.Id, time.Now(), INITIAL, time.Now(), time.Now())
	if err != nil {
		return lastRun, errors.Wrap(err, "Error inserting row")
	}
	if err := tx.Commit(); err != nil { // commmit the transaction
		return lastRun, errors.Wrap(err, "Error commiting transaction")
	}

	return lr, nil
}

func (mi repositoryImpl) UpdateExport(rs RunStatus) error {
	log.Debug("Storing", rs)

	sess, err := mi.getSession()
	if err != nil {
		log.Error("Error gettting DB session")
		return errors.Wrap(err, "Error getting conection")
	}

	tx, err := sess.Begin()
	if err != nil {
		log.Error("Error creating transaction", err)
		log.Infof("Trace %+v", err)
		return errors.Wrap(err, "Error creating transaction")
	}

	_, err = tx.Exec("UPDATE runstatus set status=?, updated_at=? where id=?", rs.Status, time.Now(), rs.Id)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return errors.Wrap(rerr, "Error performing rollback")
		}
		return errors.Wrap(err, "Error updating runstatus")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Error commiting transaction")
	}
	log.Debug("Stored run ", rs)
	return nil
}
