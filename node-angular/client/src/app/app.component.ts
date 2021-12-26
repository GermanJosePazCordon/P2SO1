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

  @ViewChild("chart1") chart!: ChartComponent;
  @ViewChild("chart2") chart2!: ChartComponent;
  @ViewChild("chart3") chart3!: ChartComponent;
  public chartOptions1: Partial<any>;
  public chartOptions2: Partial<any>;
  public chartOptions3: Partial<any>;

  constructor(private socketService : SocketService) {
    this.chartOptions1 = {
      series: [44, 55, 13, 43, 22, 45],
      chart: {
        width: 450,
        type: "pie"
      },
      labels: ["Team A", "Team B", "Team C", "Team D", "Team E", "Team F"],
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
      series: [44, 55, 13, 43, 22, 45],
      chart: {
        width: 450,
        type: "pie"
      },
      labels: ["Team A", "Team B", "Team C", "Team D", "Team E", "Team F"],
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
          data: [300, 55, 41, 67, 22, 43, 21, 33]
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
          "12-17",
          "18-24",
          "25-29",
          "30-39",
          "40-49",
          "50-59",
          "60-69",
          "70+"
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

  ngOnInit(){
    this.socketService.getNewMessage().subscribe((message: string) => {
      console.log(message)
  })
  }
}
