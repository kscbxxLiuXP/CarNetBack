package model

type Address struct {
	ID      int    `json:"id" form:"id"`
	Label   string `json:"label" form:"label"`
	Address string `json:"address" form:"address"`
}

type Job struct {
	ID        int `json:"id" form:"id"`
	VehicleID int `gorm:"column:vehicleID" json:"vehicleID" form:"vehicleID"`
	StaffID   int `gorm:"column:staffID" json:"staffID" form:"staffID"`
	TaskID    int `gorm:"column:taskID" json:"taskID" form:"taskID"`
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
}

type Task struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	MasterID    string `gorm:"column:masterID" json:"masterID" form:"masterID"`
	StartTime   string `gorm:"column:startTime" json:"startTime" form:"startTime"`
	EndTime     string `gorm:"column:endTime" json:"endTime" form:"endTime"`
	Address     string `json:"address" form:"address"`
	State       int    `json:"state" form:"state"`
	Progress    int    `json:"progress" form:"progress"`
	ExecuteTime string `gorm:"column:executeTime" json:"executeTime" form:"executeTime"`
}

type User struct {
	UserName string `gorm:"column:username;primary_key" json:"username" form:"username"`
	Password string `gorm:"column:password" json:"password" form:"password"`
}

type Vehicle struct {
	ID             int    `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Sign           string `json:"sign" form:"sign"`
	AddressID      int    `gorm:"column:addressID" json:"addressID" form:"addressID"`
	State          int    `json:"state" form:"state"`
	WorkState      int    `gorm:"column:workState" json:"workState" form:"workState"`
	RegisterTime   string `gorm:"column:registerTime" json:"registerTime" form:"registerTime"`
	ActivationTime string `gorm:"column:activationTime" json:"activationTime" form:"activationTime"`
	TaskExecuteNum int    `gorm:"column:taskExecuteNum" json:"taskExecuteNum" form:"taskExecuteNum"`
}

