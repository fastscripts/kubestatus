

// @todo in funktionen packen um doppelten Code zu vermeiden 
var memoryChartDom = document.getElementById('memoryChart');
var memoryChart = echarts.init(memoryChartDom);
var option;
chartData = JSON.parse(memoryChartDom.attributes["x-data"].value)
//console.log(chartData.node_count)


//memoryChart.showLoading();

option = {
  series: [
    {
      type: 'gauge',
      axisLine: {
        lineStyle: {
          width: 10,
          color: [
            [0.3, '#67e0e3'],
            [0.7, '#37a2da'],
            [1, '#fd666d']
          ]
        }
      },
      pointer: {
        itemStyle: {
          color: 'auto'
        }
      },
      axisTick: {
        distance: -30,
        length: 5,
        lineStyle: {
          color: '#fff',
          width: 1
        }
      },
      splitLine: {
        distance: -30,
        length: 50,
        lineStyle: {
          color: '#fff',
          width: 1
        }
      },
      axisLabel: {
        color: 'inherit',
        distance: 5,
        fontSize: 10
      },
      detail: {
        valueAnimation: true,
        formatter: '{value} %',
        color: 'inherit',
        fontSize: 15
      },
      data: [
        {
          value: 100
        }
      ]
    }
  ]
};

setInterval(function () {
  memoryChart.setOption({
    series: [
      {
        data: [
          {
           value: Math.round(chartData.memory.used * 100 / chartData.memory.capacity)
          }
        ]
      }
    ]
  });
}, 2000);

memoryChart.setOption(option);





var cpuChartDom = document.getElementById('cpuChart');
var cpuChart = echarts.init(cpuChartDom);
var option;
chartData = JSON.parse(cpuChartDom.attributes["x-data"].value)
//console.log(chartData.node_count)


//cpuChart.showLoading();

option = {
  series: [
    {
      type: 'gauge',
      axisLine: {
        lineStyle: {
          width: 8,
          color: [
            [0.3, '#5FB404'],
            [0.7, '#F7D358'],
            [1, '#DF3A01']
          ]
        }
      },
      pointer: {
        itemStyle: {
          color: 'auto'
        }
      },
      axisTick: {
        distance: -25,
        length: 5,
        lineStyle: {
          color: '#00F',
          width: 1
        }
      },
      splitLine: {
        distance: -20,
        length: 30,
        lineStyle: {
          color: '#fff',
          width: 1
        }
      },
      axisLabel: {
        color: 'inherit',
        distance: 5,
        fontSize: 10
      },
      detail: {
        valueAnimation: true,
        formatter: '{value} %',
        color: 'inherit',
        fontSize: 15
      },
      data: [
        {
          value: Math.round(chartData.cpu.used * 100 / chartData.cpu.capacity)
        }
      ]
    }
  ]
};
/*
setInterval(function () {
  cpuChart.setOption({
    series: [
      {
        data: [
          {
            value: Math.round(chartData.cpu.used * 100 / chartData.cpu.capacity)
          }
        ]
      }
    ]
  });
}, 2000);
*/
cpuChart.setOption(option);