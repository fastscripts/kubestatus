'use client'
import ReactECharts from 'echarts-for-react';

export default function BarChart({height=400}) {
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

  return <ReactECharts option={option} style={{ width: '80%',  height: `${height}px` }} />;
};


