package main

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genLine3dData() []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)
	x := []float64{2310.738003, 3551.8411, 4709.251392, 7150.631825, 7889.578078, 8261.092496, 8493.799314, 9183.754656, 9412.378906, 9873.709996, 11908.87412, 13108.13082, 14307.38751, 15549.51125, 17767.37065, 18965.6067, 20028.0971, 23552.38081, 24653.65568, 25600.81331, 26671.46886, 27528.80982, 29577.24231, 30754.04484, 32139.05874, 33556.73325, 36759.51411, 38187.39506, 39377.46595, 42875.21292, 45208.40509, 46029.00286, 47245.50844}
	y := []float64{-413.6453804, -634.9234278, -842.8580857, -1275.40662, -1407.728665, -1474.445733, -1515.587917, -1642.350191, -1681.268362, -1764.664629, -2127.160322, -2343.990579, -2561.93278, -2782.098866, -3179.064999, -3393.671325, -3581.590872, -2538.581803, -2614.194463, -2706.48628, -3149.042365, -4030.818688, -6432.630612, -7098.688644, -7422.266086, -7500.102601, -8297.370724, -8483.06637, -8520.872581, -6233.591504, -5675.392615, -5431.875558, -6663.026559}
	z := []float64{8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8869.68, 8237.22, 7597.14, 6781.8, 5280.66, 4671.06, 3878.58, 3352.8, 2788.92, 2392.68, 2263.14, 2446.02, 2621.28, 2491.74, 1333.5, 982.98, 130}
	for i := 0; i < 33; i++ {
		data = append(data, opts.Chart3DData{Value: []interface{}{
			x[i],
			y[i],
			z[i]},
		})
	}
	fmt.Println(data)
	return data
}

func line3DBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line3d example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	line3d.AddSeries("line3D", genLine3dData())
	return line3d
}

func line3DAutoRotate() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "auto rotating"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),

		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: &opts.ViewControl{
				AutoRotate: true,
			},
		}),
	)

	line3d.AddSeries("line3D", genLine3dData())
	return line3d
}

type Line3dExamples struct{}

func (Line3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		line3DBase(),
		//		line3DAutoRotate(),
	)

	f, err := os.Create("examples/html/line3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
func main() {
	a := Line3dExamples{}
	a.Examples()
}
