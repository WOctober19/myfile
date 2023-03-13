package model

type Scope string

type Role struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:100;not null;unique"`
	Scope     Scope  `json:"scope" gorm:"size:100"`
	Namespace string `json:"namespace"  gorm:"size:100"`
	Rules     Rules  `json:"rules" gorm:"type:json"`
}

type Operation string

type Rule struct {
	Resource  string    `json:"resource"`
	Operation Operation `json:"operation"`
}

type Rules []Rule
