package utils

import "fmt"
import "time"

type UtilsErr struct {

}

func (ue *UtilsErr) Error() string{
  return fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"))
}
