package pm

import "geep/module/types"

func (pm *PM) KillAll() []error {
	errors := make([]error, 0)
	for _, process := range pm.process {
		err := pm.Delete(types.DeleteMessage{Type: "delete", Name: process.name})
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
