package main1

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	dataSeed = rand.NewSource(time.Now().UnixNano())

	scatter3DColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}
)

func genScatter3dData() []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)
	x := []float64{2310.738003, 3551.8411, 4709.251392, 7150.631825, 7889.578078, 8261.092496, 8493.799314, 9183.754656, 9412.378906, 9873.709996, 11908.87412, 13108.13082, 14307.38751, 15549.51125, 17767.37065, 18965.6067, 20028.0971, 23552.38081, 24653.65568, 25600.81331, 26671.46886, 27528.80982, 29577.24231, 30754.04484, 32139.05874, 33556.73325, 36759.51411, 38187.39506, 39377.46595, 42875.21292, 45208.40509, 46029.00286, 47245.50844}
	y := []float64{-413.6453804, -634.9234278, -842.8580857, -1275.40662, -1407.728665, -1474.445733, -1515.587917, -1642.350191, -1681.268362, -1764.664629, -2127.160322, -2343.990579, -2561.93278, -2782.098866, -3179.064999, -3393.671325, -3581.590872, -2538.581803, -2614.194463, -2706.48628, -3149.042365, -4030.818688, -6432.630612, -7098.688644, -7422.266086, -7500.102601, -8297.370724, -8483.06637, -8520.872581, -6233.591504, -5675.392615, -5431.875558, -6663.026559}
	z := []float64{-413.6453804, -634.9234278, -842.8580857, -1275.40662, -1407.728665, -1474.445733, -1515.587917, -1642.350191, -1681.268362, -1764.664629, -2127.160322, -2343.990579, -2561.93278, -2782.098866, -3179.064999, -3393.671325, -3581.590872, -2538.581803, -2614.194463, -2706.48628, -3149.042365, -4030.818688, -6432.630612, -7098.688644, -7422.266086, -7500.102601, -8297.370724, -8483.06637, -8520.872581, -6233.591504, -5675.392615, -5431.875558, -6663.026559}
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

func scatter3DBase() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic Scatter3D example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddSeries("scatter3d", genScatter3dData())
	return scatter3d
}

func scatter3DDataItem() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "user-defined item style"}),
		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "MY-X-AXIS", Show: true}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "MY-Y-AXIS"}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "MY-Z-AXIS"}),
	)

	scatter3d.AddSeries("scatter3d", []opts.Chart3DData{
		{Name: "point1", Value: []interface{}{10, 10, 10}, ItemStyle: &opts.ItemStyle{Color: "green"}},
		{Name: "point2", Value: []interface{}{15, 15, 15}, ItemStyle: &opts.ItemStyle{Color: "blue"}},
		{Name: "point3", Value: []interface{}{20, 20, 20}, ItemStyle: &opts.ItemStyle{Color: "red"}},
	})

	return scatter3d
}

type Scatter3dExamples struct{}

func (Scatter3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase()
		//		scatter3DDataItem(),
	)

	f, err := os.Create("examples/html/scatter3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
func main() {
	a := Scatter3dExamples{}
	a.Examples()
}
