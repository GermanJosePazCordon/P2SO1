import { Component, ViewChild } from '@angular/core';
import { ChartComponent } from 'ng-apexcharts';
import { SocketService } from './services/socket.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'client';
  con_1: any[] = [];
  recopilados: String[] = [];
  top_3: any[] = []
  ultimos: any[] = []
  una_series: any[] = []
  una_labels: any[] = []
  completo_series: any[] = []
  completo_labels: any[] = []
  rangos: any[] = []

  @ViewChild("chart1") chart!: ChartComponent;
  @ViewChild("chart2") chart2!: ChartComponent;
  @ViewChild("chart3") chart3!: ChartComponent;
  public chartOptions1: Partial<any>;
  public chartOptions2: Partial<any>;
  public chartOptions3: Partial<any>;


  constructor(private socketService: SocketService) {
    this.chartOptions1 = {
      series: this.completo_series,
      chart: {
        width: 450,
        type: "pie",
        animations: {
          enabled: false
        },
      },
      labels: this.completo_labels,
      responsive: [
        {
          breakpoint: 480,
          options: {
            chart: {
              width: 200
            },
            legend: {
              position: "bottom"
            }
          }
        }
      ],
      title: {
        text: "Porcentaje de vacunados por departamento (una dosis)",
        align: 'center',
        margin: 10,
        offsetX: 0,
        offsetY: 0,
        floating: false,
        style: {
          fontSize: '14px',
          fontWeight: 'bold',
          fontFamily: undefined,
          color: '#263238'
        },
      },
    };

    this.chartOptions3 = {
      series: this.una_series,
      chart: {
        width: 450,
        type: "pie",
        animations: {
          enabled: false
        }
      },
      labels: this.una_labels,
      responsive: [
        {
          breakpoint: 480,
          options: {
            chart: {
              width: 200
            },
            legend: {
              position: "bottom"
            }
          }
        }
      ],
      title: {
        text: "Porcentaje de vacunados por departamento (esquema completo)",
        align: 'center',
        margin: 10,
        offsetX: 0,
        offsetY: 0,
        floating: false,
        style: {
          fontSize: '14px',
          fontWeight: 'bold',
          fontFamily: undefined,
          color: '#263238'
        },
      },
    };

    this.chartOptions2 = {
      series: [
        {
          name: "Vacunados",
          data: this.rangos
        }
      ],
      chart: {
        height: 350,
        type: "bar",
        toolbar: {
          show: false
        },
      },
      plotOptions: {
        bar: {
          columnWidth: "50%",
          endingShape: "rounded"
        }
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        width: 2
      },
      title: {
        text: "Rango de Edad de los Vacunados",
        align: 'center',
        margin: 10,
        offsetX: 0,
        offsetY: 0,
        floating: false,
        style: {
          fontSize: '14px',
          fontWeight: 'bold',
          fontFamily: undefined,
          color: '#263238'
        },
      },
      grid: {
        row: {
          colors: ["#fff", "#f2f2f2"]
        }
      },
      xaxis: {
        title: {
          text: "Rangos de edad"
        },
        labels: {
          rotate: -45
        },
        categories: [
          "0-11",
          "12-18",
          "19-26",
          "27-59",
          "60+"
        ],
        tickPlacement: "on"
      },
      yaxis: {
        title: {
          text: "Cantidad de Vacunados"
        }
      },
      fill: {
        type: "gradient",
        gradient: {
          shade: "light",
          type: "horizontal",
          shadeIntensity: 0.25,
          gradientToColors: undefined,
          inverseColors: true,
          opacityFrom: 0.85,
          opacityTo: 0.85,
          stops: [50, 0, 100]
        }
      }
    };
  }

  ngOnInit() {
    this.socketService.getNewMessage().subscribe((message: any) => {

      this.con_1 = message['con1']
      if (this.con_1 != undefined) {

        this.recopilados = []
        for (const rec of this.con_1) {
          this.recopilados.push(JSON.stringify(rec))
        }

        this.top_3 = []
        let size = message['con2'].length
        if (size > 3) {
          size = 3
        }
        for (let index = 0; index < size; index++) {
          this.top_3.push(message['con2'][index]);
        }

        this.ultimos = []
        size = message['con5'].length
        if (size > 5) {
          size = 5
        }
        for (let index = 0; index < size; index++) {
          this.ultimos.push(message['con5'][index]);
        }

        this.completo_series = []
        this.completo_labels = []
        for (let dato of message['con4']) {
          this.completo_series.push(dato['total'])
          this.completo_labels.push(dato['_id'])
        }

        this.chartOptions3['series'] = this.completo_series
        this.chartOptions3['labels'] = this.completo_labels

        this.una_series = []
        this.una_labels = []
        for (let dato of message['con3']) {
          this.una_series.push(dato['total'])
          this.una_labels.push(dato['_id'])
        }

        this.chartOptions1['series'] = this.una_series
        this.chartOptions1['labels'] = this.una_labels

        let rango = ''
        this.rangos = []
        for (let i = 1; i < 6; i++) {
          rango = "rango" + i
          this.rangos.push(message['con6'][rango])
        }

        this.chartOptions2['series'] = [{
          name: "vacunados",
          data: this.rangos
        }];
      }
    })
  }
}
