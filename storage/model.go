package storage


type Model struct {}


var (
	TimeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)
//
//// UnmarshalJSON implements json unmarshal interface.
//func (t *time.Time) UnmarshalJSON(data []byte) (err error) {
//	now, err := time.ParseInLocation(`"`+TimeFormart+`"`, string(data), time.Local)
//	*t = time.Time(now)
//	return
//}
//
//// MarshalJSON implements json marshal interface.
//func (t time.Time) MarshalJSON() ([]byte, error) {
//	b := make([]byte, 0, len(TimeFormart)+2)
//	b = append(b, '"')
//	b = time.Time(t).AppendFormat(b, TimeFormart)
//	b = append(b, '"')
//	return b, nil
//}
//
//func (t time.Time) String() string {
//	return time.Time(t).Format(TimeFormart)
//}
//
//func (t time.Time) local() time.Time {
//	loc, _ := time.LoadLocation(zone)
//	return time.Time(t).In(loc)
//}
//
//func (t time.Time) Value() (driver.Value, error) {
//	var zeroTime time.Time
//	var ti = time.Time(t)
//	if ti.UnixNano() == zeroTime.UnixNano() {
//		return "0000-00-00 00:00:00", nil	//Time的默认值, 写
//	}
//	return ti, nil
//}
//
//func (t *time.Time) Scan(v interface{}) error {
//	value, ok := v.(time.Time)
//	if ok {
//		*t = time.Time(value)
//		return nil
//	}
//	return fmt.Errorf("can not convert %v to timestamp", v)
//}

func TimeZero() string{
	return "0000-00-00 00:00:00"
}
