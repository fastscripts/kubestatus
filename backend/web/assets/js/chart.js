


if (document.getElementById('memoryChart')) {

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



}

if (document.getElementById('cpuChart')) {

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
}


if (document.getElementById('usageChart')){


var chartElement = document.getElementById('usageChart');
var usageCahrt = echarts.init(chartElement);
const formatUtil = echarts.format;



usageCahrt.showLoading();


$.get('/data/usage.tree.json', function (usageData) {
  usageCahrt.hideLoading();



  /*  fetch('/data/disk.tree.json').then(function (response) {
    return response.json();
  }).then(function (usageData) {
    usageCahrt.hideLoading();
  }).catch(function (err) {
    ;console.log('Fetch Error :-S', err);
  }); */


  function getLevelOption() {
    return [
      {
        itemStyle: {
          borderColor: '#777',
          borderWidth: 0,
          gapWidth: 1
        },
        upperLabel: {
          show: false
        }
      },
      {
        itemStyle: {
          borderColor: '#555',
          borderWidth: 5,
          gapWidth: 1
        },
        emphasis: {
          itemStyle: {
            borderColor: '#ddd'
          }
        }
      },
      {
        colorSaturation: [0.35, 0.5],
        itemStyle: {
          gapWidth: 1,
          borderColorSaturation: 0.6
        }
      }
    ];
  }



    var option = {
      title: {
        text: 'Usage',
        left: 'center'
      },
      tooltip: {
        formatter: function (info) {
          info.value = info.data.value;
          var value = info.value;
          var cpu = info.data.cpu;
          var treePathInfo = info.treePathInfo;
          var treePath = [];
          for (var i = 1; i < treePathInfo.length; i++) {
            treePath.push(treePathInfo[i].name);
          }
          return [
            '<div class="tooltip-title">' +
              echarts.format.encodeHTML(treePath.join('/')) +
              '</div>',
            'Ram Usage: ' + value + ' KB <br>',
            'CPU Usage: '+ cpu + 'VCPU'

          ].join('');
        }
      },
      series: [
        {
          name: 'CPU Usage',
          type: 'treemap',
          visibleMin: 300,
          label: {
            show: true,
            formatter: '{b}'
          },
          itemStyle: {
            borderColor: '#fff'
          },
          levels: getLevelOption(),
          data: usageData
        }
      ],
      color: ['#37A2DA', '#32C5E9', '#67E0E3', '#9FE6B8', '#FFDB5C', '#ff9f7f']

    }

  // Display the chart using the configuration items and data just specified.
  usageCahrt.setOption(option);

  })
}