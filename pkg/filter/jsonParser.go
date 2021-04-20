package filter

import (
	"CS4227/pkg/factory"
	"encoding/json"
	"log"
	"strings"
)

type cpu struct {
	Id           int
	Manufacturer string
	Name         string
	Cores        int
	Clock        string
	Tdp          int
	Socket       string
	Price        int
}

type drive struct {
	Id           int
	Name         string
	Manufacturer string
	Price        int
	Size         int
	Tdp          int
	Technology   string
}

type gpu struct {
	Id           int
	Name         string
	Manufacturer string
	Price        int
	Memory       int
	Clock        string
	Tdp          int
}

type mobo struct {
	Id           int
	Manufacturer string
	Name         string
	Price        int
	Socket       string
	Tdp          int
	RamSlots     int
	MaxRam       int
	DriveSlots   int
}

type psu struct {
	Id           int
	Name         string
	Manufacturer string
	Price        int
	Rating       string
	Capacity     int
}

type ram struct {
	Id           int
	Name         string
	Manufacturer string
	Price        int
	Clock        string
	Memory       int
}

func JsonParse(request string, data string) []factory.Component {
	asusFactory, _ := factory.GetPartsFactory("asus")
	amdFactory, _ := factory.GetPartsFactory("amd")
	seagateFactory, _ := factory.GetPartsFactory("seagate")
	nvidiaFactory, _ := factory.GetPartsFactory("nvidia")
	intelFactory, _ := factory.GetPartsFactory("intel")

	data = strings.Replace(data, "{[", "[", 1)
	data = strings.Replace(data, "]}", "]", 1)

	var returned []factory.Component

	if request == "cpu" {
		var result []cpu

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakeCPU()
				made.SetData(object.Id, object.Name, object.Price, object.Cores, object.Clock, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakeCPU()
				made.SetData(object.Id, object.Name, object.Price, object.Cores, object.Clock, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakeCPU()
				made.SetData(object.Id, object.Name, object.Price, object.Cores, object.Clock, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakeCPU()
				made.SetData(object.Id, object.Name, object.Price, object.Cores, object.Clock, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakeCPU()
				made.SetData(object.Id, object.Name, object.Price, object.Cores, object.Clock, object.Socket, object.Tdp)
				returned = append(returned, made)
			}
		}
	}

	if request == "gpu" {
		var result []gpu

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakeGPU()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakeGPU()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakeGPU()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakeGPU()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakeGPU()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Memory, object.Clock)
				returned = append(returned, made)
			}
		}
	}

	if request == "drive" {
		var result []drive

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakeDRIVE()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Size, object.Technology)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakeDRIVE()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Size, object.Technology)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakeDRIVE()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Size, object.Technology)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakeDRIVE()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Size, object.Technology)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakeDRIVE()
				made.SetData(object.Id, object.Name, object.Price, object.Tdp, object.Size, object.Technology)
				returned = append(returned, made)
			}
		}
	}

	if request == "mbd" {
		var result []mobo

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakeMOBO()
				made.SetData(object.Id, object.Name, object.Price, object.RamSlots, object.MaxRam, object.DriveSlots, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakeMOBO()
				made.SetData(object.Id, object.Name, object.Price, object.RamSlots, object.MaxRam, object.DriveSlots, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakeMOBO()
				made.SetData(object.Id, object.Name, object.Price, object.RamSlots, object.MaxRam, object.DriveSlots, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakeMOBO()
				made.SetData(object.Id, object.Name, object.Price, object.RamSlots, object.MaxRam, object.DriveSlots, object.Socket, object.Tdp)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakeMOBO()
				made.SetData(object.Id, object.Name, object.Price, object.RamSlots, object.MaxRam, object.DriveSlots, object.Socket, object.Tdp)
				returned = append(returned, made)
			}
		}
	}

	if request == "psu" {
		var result []psu

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakePSU()
				made.SetData(object.Id, object.Name, object.Price, object.Capacity, object.Rating)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakePSU()
				made.SetData(object.Id, object.Name, object.Price, object.Capacity, object.Rating)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakePSU()
				made.SetData(object.Id, object.Name, object.Price, object.Capacity, object.Rating)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakePSU()
				made.SetData(object.Id, object.Name, object.Price, object.Capacity, object.Rating)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakePSU()
				made.SetData(object.Id, object.Name, object.Price, object.Capacity, object.Rating)
				returned = append(returned, made)
			}
		}
	}

	if request == "ram" {
		var result []ram

		jsonErr := json.Unmarshal([]byte(data), &result)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range result {
			object := result[i]
			manufacturer := strings.ToLower(object.Manufacturer)
			if manufacturer == "asus" {
				made := asusFactory.MakeRAM()
				made.SetData(object.Id, object.Name, object.Price, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "amd" {
				made := amdFactory.MakeRAM()
				made.SetData(object.Id, object.Name, object.Price, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "intel" {
				made := intelFactory.MakeRAM()
				made.SetData(object.Id, object.Name, object.Price, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "nvidia" {
				made := nvidiaFactory.MakeRAM()
				made.SetData(object.Id, object.Name, object.Price, object.Memory, object.Clock)
				returned = append(returned, made)
			}

			if manufacturer == "seagate" {
				made := seagateFactory.MakeRAM()
				made.SetData(object.Id, object.Name, object.Price, object.Memory, object.Clock)
				returned = append(returned, made)
			}
		}
	}

	return returned
}
