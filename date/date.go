package date

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"time"
)

type Date struct {
	time   time.Time
	inited bool
}

var (
	_ sql.Scanner   = (*Date)(nil)
	_ driver.Valuer = (*Date)(nil)
)

func (d *Date) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case []byte:
		if len(v) == 0 {
			return nil
		}
		vd, err := Parse("2006-01-02", string(v))
		*d = vd
		return err
	case string:
		if v == "" {
			return nil
		}
		vd, err := Parse("2006-01-02", v)
		*d = vd
		return err
	case time.Time:
		*d = NewTime(v)
		return nil
	case *time.Time:
		*d = NewTime(*v)
		return nil
	case Date:
		*d = v
		return nil
	case *Date:
		*d = *v
		return nil
	}

	return fmt.Errorf("date: unknown src type %T", src)
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func Now() Date {
	return NewTime(time.Now())
}

func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return NewTime(t), nil
}

func ParseInLocation(layout, value string, loc *time.Location) (Date, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return Date{}, err
	}
	return NewTime(t), nil
}

func Unix(sec int64, nsec int64) Date {
	return NewTime(time.Unix(sec, nsec))
}

func New(year int, month time.Month, day int, loc *time.Location) Date {
	var d Date
	d.time = time.Date(year, month, day, 0, 0, 0, 0, loc)
	d.inited = true
	return d
}

func NewTime(t time.Time) Date {
	var d Date
	d.time = t
	d.init()
	return d
}

func (d *Date) init() {
	if d.inited {
		return
	}
	d.time = time.Date(d.time.Year(), d.time.Month(), d.time.Day(), 0, 0, 0, 0, d.time.Location())
	d.inited = true
}

func (d Date) Add(dur time.Duration) Date {
	return NewTime(d.time.Add(dur))
}

func (d Date) AddDate(years int, months int, days int) Date {
	return NewTime(d.time.AddDate(years, months, days))
}

func (d Date) After(u Date) bool {
	return d.time.After(u.time)
}

func (d Date) AppendFormat(b []byte, layout string) []byte {
	return d.time.AppendFormat(b, layout)
}

func (d Date) Before(u Date) bool {
	return d.time.Before(u.time)
}

func (d Date) Clock() (hour, min, sec int) {
	return d.time.Clock()
}

func (d Date) Date() (year int, month time.Month, day int) {
	return d.time.Date()
}

func (d Date) Day() int {
	return d.time.Day()
}

func (d Date) Equal(u Date) bool {
	return d.time.Equal(u.time)
}

func (d Date) Format(layout string) string {
	return d.time.Format(layout)
}

func (d *Date) GobDecode(data []byte) error {
	if err := d.time.GobDecode(data); err != nil {
		return err
	}
	d.init()
	return nil
}

func (d Date) GobEncode() ([]byte, error) {
	return d.time.GobEncode()
}

func (d Date) Hour() int {
	return d.time.Hour()
}

func (d Date) ISOWeek() (year, week int) {
	return d.time.ISOWeek()
}

func (d Date) In(loc *time.Location) Date {
	return NewTime(d.time.In(loc))
}

func (d Date) IsZero() bool {
	return d.time.IsZero()
}

func (d Date) Local() Date {
	return NewTime(d.time.Local())
}

func (d Date) Location() *time.Location {
	return d.time.Location()
}

func (d Date) MarshalBinary() ([]byte, error) {
	return d.time.MarshalBinary()
}

func (d Date) MarshalJSON() ([]byte, error) {
	return d.time.MarshalJSON()
}

func (d Date) MarshalText() ([]byte, error) {
	return d.time.MarshalText()
}

func (d Date) Minute() int {
	return d.time.Minute()
}

func (d Date) Month() time.Month {
	return d.time.Month()
}

func (d Date) Nanosecond() int {
	return d.time.Nanosecond()
}

func (d Date) Round(dur time.Duration) Date {
	return NewTime(d.time.Round(dur))
}

func (d Date) Second() int {
	return d.time.Second()
}

func (d Date) String() string {
	return d.time.Format("2006-01-02")
}

func (d Date) Sub(u Date) time.Duration {
	return d.time.Sub(u.time)
}

func (d Date) Truncate(dur time.Duration) Date {
	return NewTime(d.time.Truncate(dur))
}

func (d Date) UTC() Date {
	return NewTime(d.time.UTC())
}

func (d Date) Unix() int64 {
	return d.time.Unix()
}

func (d Date) UnixNano() int64 {
	return d.time.UnixNano()
}

func (d *Date) UnmarshalBinary(data []byte) error {
	if err := d.time.UnmarshalBinary(data); err != nil {
		return err
	}
	d.init()
	return nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if err := d.time.UnmarshalJSON(data); err != nil {
		return err
	}
	d.init()
	return nil
}

func (d *Date) UnmarshalText(data []byte) error {
	if err := d.time.UnmarshalText(data); err != nil {
		return err
	}
	d.init()
	return nil
}

func (d Date) Weekday() time.Weekday {
	return d.time.Weekday()
}

func (d Date) Year() int {
	return d.time.Year()
}

func (d Date) YearDay() int {
	return d.time.YearDay()
}

func (d Date) Zone() (name string, offset int) {
	return d.time.Zone()
}
