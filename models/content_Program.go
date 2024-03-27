package models

import (
	"time"

	"gorm.io/gorm"
)

type Program struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	IsGenerated bool   `json:"is_generated"`
	Name        string `json:"name"`
	Meaning     string `json:"meaning"`
	Guidance    string `json:"guidance"`
	BatchId     string `json:"batch_id"`
	UserId      uint32 `json:"user_id"`
	UserKey     string `json:"user_key"`
	Desc        string `json:"desc"`
	Longdesc    string `json:"longdesc"`
	Difficulty  string `json:"difficulty"`
	NbViews     uint32 `json:"nb_views"`

	//POPULATED
	//LinksProgramWorkout []LinkProgramWorkout `gorm:"foreignKey:program_id" json:"links,omitempty"`
	Plangroups []Plangroup `gorm:"foreignKey:program_id" json:"plangroups,omitempty"`
}

func (c *Program) GetId() uint32 {
	return c.ID
}

func (c *Program) TableName() string {
	return "api_content.program"
}

func (c *Program) GetNextWorkout(currentWorkoutId uint32) (bool, *Plangroup, *Workout, *Workout) {
	//plangroupIdx := -1
	//planIdx := -1
	isSamePlangroup := false
	nextPlanIdx := -1
	nextPlangroupIdx := -1
	currentWorkout := Workout{}
	firstWorkoutOfGroup := Workout{}
	Nplangroups := len(c.Plangroups)
	nbPlanPerGroup := 5
	for i, plangroup := range c.Plangroups {
		//	_ := len(plangroup.Plans)
		for j, plan := range plangroup.Plans {

			workout := plan.Workout
			if workout != nil && workout.ID == currentWorkoutId {
				//			plangroupIdx = i
				//			planIdx = j
				currentWorkout = *workout
				firstWorkoutOfGroup = *(plangroup.Plans[0].Workout)
				if j+1 < nbPlanPerGroup {
					nextPlanIdx = j + 1
					nextPlangroupIdx = i
					isSamePlangroup = true
					break
				} else if (j+1) == nbPlanPerGroup && (i+1) < Nplangroups {
					nextPlanIdx = 0
					nextPlangroupIdx = i + 1
					isSamePlangroup = false
					break
				} else { //error case, where too many plans per group
					nextPlanIdx = 0
					nextPlangroupIdx = i + 1
					isSamePlangroup = false
				}
			}
		}
	}
	if nextPlanIdx != -1 && nextPlangroupIdx != -1 {
		pPlangroup := &c.Plangroups[nextPlangroupIdx]
		//	pWorkout := c.Plangroups[nextPlangroupIdx].Plans[nextPlanIdx].Workout
		return isSamePlangroup, pPlangroup, &currentWorkout, &firstWorkoutOfGroup
	}
	return false, nil, nil, nil
}
