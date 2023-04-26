package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

var x = [10]string{
	"<b>LADA 4x4 - UPDATED NIVA</b><br>AvtoVAZ has officially started production of the updated Lada 4x4. Niva received a side airbag, electric and heated mirrors, more comfortable seats from the Togliatti Lada-Plast company with a different stuffing and pronounced lateral support, another front bumper with fog lights integrated into it, as well as wheels with a common tire size - 175/80 R16. Also, the updated SUV received a new steering wheel with a front airbag and a steering column casing. Sales of Lada 4x4 FL should begin in January - February 2020. Officially, the price of the car has not yet been announced.",
	"<b>GEELY FY11 - CROSS-COUPE ON THE VOLVO PLATFORM</b><br>Geely FY11 has a rich equipment. In the list of options there is a driver recognition function by face, which identifies the person behind the wheel and activates previously made settings. There is adaptive cruise control, automatic emergency braking, lane keeping and collision avoidance systems, blind spot monitoring, and an all-round camera. Geely FY11 received Volvo engines: a 1.5-liter turbo engine with a capacity of 177 hp and a 2-liter turbo engine with a capacity of 238 hp, combined with a 7-speed robot and an 8-band Aisin automatic. The cost of FY11 in China is from 132.8 thousand yuan.",
	"<b>KIA SELTOS - KILLER OF HYUNDAI CRETA</b><br>The premiere of the compact crossover Kia Seltos took place at the end of November in China. The Russian version will be assembled at the Kaliningrad Avtotor plant. Seltos with three engines will be available. These are 1.6-liter and 2-liter atmospheric engines with a capacity of 121-123 and 149 hp, respectively. The basic engine is combined with a 6-speed manual transmission, as well as with a classic automatic. The top version is paired with a continuously variable variator. There is a front independent spring suspension with a stabilizer bar and a semi-independent rear; disc brakes, steering wheel with electric power. The company has already announced the cost of the SUV for Russia.",
	"<b>PEUGEOT 2008 - A NEW GENERATION OF COMPACT CROSSOVER</b><br>Peugeot 2008 is expected on the Russian market in the spring of 2020. The car is built on a new modular CMP platform, its length is 4300 mm, the wheelbase is 2600 mm, the volume of luggage space is 434 liters. Features - two-tone body, original running lights connected to the front optics, wide black radiator grille, three-dimensional wheel arches and 18-inch alloy wheels with double spokes. The crossover is equipped with a 1.2-liter gasoline turbo engine PureTech power. 100, 130 or 155 hp, as well as a 1.5-liter BlueHDi turbodiesel for 100 or 130 hp. Power units are paired with a 6-speed manual transmission, or with an 8-speed automatic.",
	"<b>CHERY EXEED TXL - DESIGNER PREMIUM CROSSOVER</b><br>The Chery Exeed TXL, designed by former BMW stylist Kevin Rice, will reach Russia in the third quarter of 2020. The car is built on the Chery Tiggo 8 platform and is equipped with the same 1.6-liter turbocharged engine with a capacity of 197 hp paired with a 7-band robot with two clutches. The letter L in the name of the crossover indicates that it is elongated - 4,775 mm. The car will receive a fairly rich basic version: dual-zone climate control, a full winter package, 6 airbags, keyless entry and button start, leather interior, front fog lights. The cost in China is from 147.9 to 175.9 thousand yuan.",
	"<b>SKODA OCTAVIA - THE 4TH GENERATION OF THE CZECH BESTSELLER</b><br>Serial production of the 4th generation of the brand began at the plant in the Czech Mlada-Boleslav at the end of November. Octavia will be available with a gasoline 2-liter TSI turbo engine with a capacity of 190 hp, 2.0-liter turbodiesels with a return of 115, 150 and 200 hp, as well as with a plug-in hybrid powerplant, a mild hybrid and a CNG gas version. There is a digital dashboard with a 10-inch multimedia system screen and the ability to project information onto the windshield; three-zone climate control, massage seats, a new two-spoke steering wheel. According to the results of the EuroNCAP safety test, the overall rating was the highest 5 stars. The new generation will reach the domestic market only at the end of 2020.",
	"<b>VOLKSWAGEN JETTA - SEVENTH-GENERATION SEDAN</b><br>Volkswagen Jetta will be produced in Mexico. The car received a fresh design, it was built on a modular MQB platform with a transverse engine arrangement. The Jetta will be delivered to Russia with two petrol engines of 1.4 and 1.6 liters, with a capacity of 110 and 150 hp and a maximum torque of 155 and 250 Nm. Paired with a 100-horsepower engine, a 5-speed manual or a 6-speed automatic transmission will go. Jetta received a new dashboard, an upgraded multimedia system with support for Apple CarPlay, Android Auto and MirrorLink, an electronic stability control system, emergency braking, tire pressure monitoring.",
	"<b>MITSUBISHI PAJERO - THE REDESIGNED PAJERO SPORT</b><br>The SUV received the Dynamic Shield corporate identity. The front two-story optics appeared, the taillights in the lower part became slightly shorter, a spoiler appeared on the rear trunk door. The analog dashboard was replaced by an eight-inch screen with LED indicators of fuel level and coolant temperature. A small glove compartment appeared in the central tunnel, the control unit of the climate system changed. The Pajero Sport is equipped with a 2.4-liter turbodiesel with 181 hp and an 8-speed automatic transmission. A gasoline 3-liter V6 with 209 hp will also be available in Russia. with the same gearbox and all-wheel drive. The car should appear on the Russian market in the second half of 2020.",
	"<b>HYUNDAI PALISADE - THE FLAGSHIP CROSSOVER</b><br>Hyundai Palisade has already gone on sale in Korea, the USA and Canada, and in the second half of 2020 it will reach Russia. The SUV has impressive dimensions - 4981 x 1976 x 1750 mm and accommodates seven people. Buyers will be offered a choice of two powertrains: a 3.8-liter gasoline V6 with a capacity of 295 hp and a 150-horsepower 2.2-liter diesel. Both are paired with an 8-band automatic. The peculiarity of the car is the driving modes, which have been replenished with non-standard ones - sand and snow. Perforated leather, soft plastic, carbon fiber inserts are used in the interior decoration. The digital instrument panel has a 12.3-inch screen, information from the speedometer or navigator will be displayed by the projector on the windshield.",
	"<b>AUDI E-TRON - FUTURISTIC CAR</b><br>The Audi E-tron electric car is made on the MLB Evo modular platform. The company's first all-electric SUV is equipped with asynchronous motors that can maintain a maximum power of 360 hp. The car's power reserve is about 400 kilometers. This is the first mass-produced passenger electric car in the world equipped with an integrated electrohydraulic brake control system. A coupe-like crossover Audi e-tron Sportback, made on the same platform, will also arrive in Russia. The basic version of the model is driven by two electric motors, which give a total of 313 hp. The cost of a car in Europe is 71,350 euros. In Russia, it is expected in the second quarter of 2020.",
}

func main() {
	http.HandleFunc("/", pict)
	fmt.Println("Listen and serve on port :8080")
	http.ListenAndServe(":8080", nil)
}

func pict(w http.ResponseWriter, r *http.Request) {
	var k int
	if k, _ = strconv.Atoi(r.FormValue("pic")); k > 0 {
		f, err := os.Open("static/img" + r.FormValue("pic") + ".jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		img, _, _ := image.Decode(f)
		var buff bytes.Buffer
		opt := jpeg.Options{
			Quality: 90,
		}
		err = jpeg.Encode(&buff, img, &opt)
		if err != nil {
			panic(err)
		}
		encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
		templ, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := templ.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write([]byte("<p><img src=\"data:image/jpeg;base64," + encodedString + "\">" + x[k-1] + "</p>"))
	} else {
		templ, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := templ.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
