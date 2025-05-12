/*
Crie um novo tipo: veículo
	O tipo subjacente deve ser struct
	Deve conter os campos: portas, cor
Crie dois novos tipos: caminhonete e sedan
	Os tipos subjacentes devem ser struct
	Ambos devem conter "veículo" como struct embutido
	O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
	O tipo sedan deve conter um campo bool chamado "modeloLuxo"
Usando os structs veículo, caminhonete e sedan:
	Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
	Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
Demonstre estes valores.
Demonstre um único campo de cada um dos dois.
*/

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle          vehicle
	four_wheel_drive bool
}

type sedan struct {
	vehicle      vehicle
	luxury_model bool
}

func main() {
	car1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		four_wheel_drive: true,
	}

	car2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury_model: false,
	}

	fmt.Println(car1)
	fmt.Println(car1.vehicle.color)

	fmt.Println(car2)
	fmt.Println(car2.vehicle.color)
}
