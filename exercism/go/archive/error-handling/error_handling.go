// Package erratum defines some error handling tasks
package erratum

// testVersion is the suite we're up against
const testVersion = 2

// Use receives a ResourceOpener and handles the open/close
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

TryOpen:
	for {
		r, err = o()
		switch err.(type) {
		case TransientError:
			continue
		case nil:
			break TryOpen
		default:
			return err
		}
	}
	defer r.Close()

	errChan := make(chan error)
	go func(r Resource) {
		defer func() {
			err := recover()
			switch err.(type) {
			case FrobError:
				r.Defrob(err.(FrobError).defrobTag)
				errChan <- err.(FrobError).inner
			case error:
				errChan <- err.(error)
			default:
				errChan <- nil
			}
		}()
		r.Frob(input)
	}(r)

	return <-errChan
}
