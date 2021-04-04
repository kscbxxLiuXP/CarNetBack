package model

type Address struct {
	ID      int    `json:"id" form:"id"`
	Label   string `json:"label" form:"label"`
	Address string `json:"address" form:"address"`
}

type Job struct {
	ID        int    `json:"id" form:"id"`
	VehicleID int    `gorm:"column:vehicleID" json:"vehicleID" form:"vehicleID"`
	StaffID   int    `gorm:"column:staffID" json:"staffID" form:"staffID"`
	Step      int    `gorm:"column:step" json:"step" form:"step"`
	Step1Time string `gorm:"column:step1Time" json:"step1Time" form:"step1Time"`
	Step2Time string `gorm:"column:step2Time" json:"step2Time" form:"step2Time"`
	Step3Time string `gorm:"column:step3Time" json:"step3Time" form:"step3Time"`
	Step4Time string `gorm:"column:step4Time" json:"step4Time" form:"step4Time"`
}

type Log struct {
	ID        int    `json:"id" form:"id"`
	VehicleID int    `gorm:"column:vehicleID" json:"vehicleID" form:"vehicleID"`
	StaffID   int    `gorm:"column:staffID" json:"staffID" form:"staffID"`
	Time      string `json:"time" form:"time"`
}

type Permission struct {
	ID        int    `json:"id" form:"id"`
	VehicleID int    `gorm:"column:vehicleID" json:"vehicleID" form:"vehicleID"`
	StaffID   int    `gorm:"column:staffID" json:"staffID" form:"staffID"`
	Time      string `json:"time" form:"time"`
}

type Setting struct {
	ID    int `json:"id" form:"id"`
	Value int `json:"value" form:"value"`
}

type Staff struct {
	ID        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	IDNumber  string `gorm:"column:idNumber" json:"idNumber" form:"idNumber"`
	Gender    string `json:"gender" form:"gender"`
	Age       int    `json:"age" form:"age"`
	Photoed   int    `json:"photoed" form:"photoed"`
	State     int    `json:"state" form:"state"`
	WorkState int    `gorm:"column:workState" json:"workState" form:"workState"`
	TaskNum   int    `gorm:"column:taskNum" json:"taskNum" form:"taskNum"`
}

type Task struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	StartTime   string `gorm:"column:startTime" json:"startTime" form:"startTime"`
	EndTime     string `gorm:"column:endTime" json:"endTime" form:"endTime"`
	Address     int    `json:"address" form:"address"`
}

type User struct {
	UserName string `gorm:"column:username;primary_key" json:"username" form:"username"`
	Password string `gorm:"column:password" json:"password" form:"password"`
}

type Vehicle struct {
	ID             int    `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Sign           string `json:"sign" form:"sign"`
	Identification string `json:"identification" form:"identification"`
	AddressID      int    `gorm:"column:addressID" json:"addressID" form:"addressID"`
	State          int    `json:"state" form:"state"`
	WorkState      int    `gorm:"column:workState" json:"workState" form:"workState"`
	RegisterTime   string `gorm:"column:registerTime" json:"registerTime" form:"registerTime"`
	ActivationTime string `gorm:"column:activationTime" json:"activationTime" form:"activationTime"`
	TaskExecuteNum int    `gorm:"column:taskExecuteNum" json:"taskExecuteNum" form:"taskExecuteNum"`
	WorkStaffID    int    `gorm:"column:workStaffID" json:"workStaffID" form:"workStaffID"`
}

type MaxStruct struct {
	Max int `json:"max"`
}

type CountStruct struct {
	Count int `json:"count"`
}
