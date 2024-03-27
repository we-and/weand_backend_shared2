package models

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

/*
type DurationMap struct {
	Durations map[int]int `json:"durations"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *DurationMap) Scan(value interface{}) error {
	//value = {"durations":{15:12}}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	//	result := json.RawMessage{}

	err := json.Unmarshal(bytes, &j)
	//	d:=result//["durations"]
	//dur := d.(map[int]int)

	//	j.Durations=d
	//	for key, value := range dur {
	// Each value is an `any` type, that is type asserted as a string
	//	  fmt.Println(key, value.(string))
	//	}
	//*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j DurationMap) Value() (driver.Value, error) {
	if len(j.Durations) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}*/

type Workout struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	BatchId string `json:"batch_id"`

	IsGenerated bool              `json:"is_generated"`
	Name        string            `json:"name"`
	Desc        string            `json:"desc"`
	NbViews     uint32            `json:"nb_views"`
	NbMoves     uint32            `json:"nb_moves"`
	Durations   datatypes.JSONMap `json:"durations"`
	//	DurationsMap map[int]int `json:"durations_map"`

	//POPULATED
	//LinksProgramWorkout []LinkProgramWorkout `gorm:"foreignKey:workout_id",json:"links,omitempty"`
	LinksTagWorkout       []LinkTagWorkout       `gorm:"foreignKey:workout_id",json:"linktags,omitempty"`
	LinksMovegroupWorkout []LinkMovegroupWorkout `gorm:"foreignKey:workout_id",json:"linkmoves,omitempty"`
}

func (c *Workout) GetId() uint32 {
	return c.ID
}

func (c *Workout) GetMoveList() []Move {
	res := []Move{}
	for _, mg := range c.LinksMovegroupWorkout {
		if mg.Movegroup != nil {
			g := *mg.Movegroup
			for _, gf := range g.LinksMovegroupMove {
				if gf.Move != nil {
					m := *gf.Move
					res = append(res, m)
				}
			}
		}
	}
	return res
}

func (c *Workout) TableName() string {
	return "api_content.workout"
}

func (c *Workout) GetMoveCount() int {
	n := 0
	for _, mg := range c.LinksMovegroupWorkout {
		if mg.Movegroup != nil {
			g := *mg.Movegroup
			for _, gf := range g.LinksMovegroupMove {
				if gf.Move != nil {
					m := *gf.Move
					if m.IsChiral {
						n = n + 1

					} else {
						n = n + 1
					}
				}
			}
		}
	}
	return n
}

func (c *Workout) GetDurationsMap() map[string]interface{} {
	res := map[string]interface{}{}
	durations := []int{3, 15, 20, 30, 60, 90, 120, 180}
	for _, duration := range durations {
		str := fmt.Sprintf("%d", duration)
		res[str] = c.GetDuration(duration)
	}
	return res
}
func (c *Workout) GetDuration(durationPerMoveSec int) int {
	n := 0
	for _, mg := range c.LinksMovegroupWorkout {
		if mg.Movegroup != nil {
			g := *mg.Movegroup
			for _, gf := range g.LinksMovegroupMove {
				if gf.Move != nil {
					m := *gf.Move
					if m.IsChiral {
						n = n + durationPerMoveSec

					} else {

						n = n + 1*durationPerMoveSec
					}
				}
			}
		}
	}
	return n
}
