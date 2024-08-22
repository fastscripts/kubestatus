import React from 'react';
import ReactECharts from 'echarts-for-react';

const MyChart = () => {
  const option = {
    title: {
      text: 'ECharts Example',
    },
    tooltip: {},
    xAxis: {
      data: ['category1', 'category2', 'category3', 'category4', 'category5', 'category6'],
    },
    yAxis: {},
    series: [
      {
        name: 'Sales',
        type: 'bar',
        data: [5, 20, 36, 10, 10, 20],
      },
    ],
  };

  return <ReactECharts option={option} style={{ height: 400 }} />;
};

export default MyChart;
