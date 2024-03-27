package models

/*
type LinkProgramWorkout struct {
	ID        uint32         `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	ProgramId uint32 `json:"program_id"`
	WorkoutId  uint32 `json:"workout_id"`

	//POPULATED
	Program *Program `gorm:"foreignKey:program_id" json:"program,omitempty"`
	Workout *Workout `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *LinkProgramWorkout) GetId() uint32 {
	return c.ID
}

func (c *LinkProgramWorkout) TableName() string {
	return "api_content.link_program_workout"
}
*/