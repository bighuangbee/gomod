package storage

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Model struct {
	ID uint `gorm:"column:id;primary_key" json:"id"`
}

type BaseModel struct {
	ID        uint `gorm:"column:id;primary_key" json:"id"`
	CreatedAt Time `gorm:"column:create_time" json:"create_time"`
	UpdatedAt Time `gorm:"column:update_time" json:"update_time"`
	DeletedAt Time `gorm:"column:delete_time" json:"create_time";sql:"index"`
}

type Time time.Time

var (
	TimeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

// UnmarshalJSON implements json unmarshal interface.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON implements json marshal interface.
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormart)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return "0000-00-00 00:00:00", nil	//Time的默认值, Writor
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func TimeZero() string{
	return "0000-00-00 00:00:00"
}
