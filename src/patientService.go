package src

import "errors"

type PatientService struct {
	PatientMap map[string]*Patient
}

func (patientService *PatientService) AddPatient(name string) error {
	if patientService != nil {
		patientService.PatientMap[name] = &Patient{
			Name: name,
		}

		return nil
	}

	return errors.New("patient service unavailable")
}
