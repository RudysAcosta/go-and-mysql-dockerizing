package serializers

import (
	"encoding/json"

	"github.com/rudysacosta/go-and-mysql-dockerizing/api/models"
)

func SerializeRol(rol models.Rol) ([]byte, error) {
	jsonData, err := json.Marshal(rol)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
