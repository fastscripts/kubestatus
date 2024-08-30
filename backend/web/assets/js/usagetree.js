document.addEventListener("DOMContentLoaded", () => {
  fetch("/usage/data/memory")
    .then((res) => res.json())
    .then((res) => {
      drawTreeMap(res);
    });

  divToolTip = d3
    .select("body")
    .append("div")
    .attr("class", "tooltip")
    .style("opacity", 0);
});

const elem = document.getElementsByTagName("body");

const bodyColor = window
  .getComputedStyle(elem[0], null)
  .getPropertyValue("background-color");

d3.select("#Memory").on("click", setType("mem"));
d3.select("#Cpu").on("click", setType("cpu"));
d3.select("#memRequested").on("click", setType("memRequested"));
d3.select("#cpuRequested").on("click", setType("cpuRequested"));
d3.select("#memLimit").on("click", setType("memLimit"));
d3.select("#cpuLimit").on("click", setType("cpuLimit"));

var SearchType = "mem";

function setType(stype) {
  return function () {
    SearchType = stype;

    if (stype === "mem") {
      SearchType = "mem";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/usage/data/memory")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    } else if (stype === "cpu") {
      SearchType = "cpu";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/usage/data/memory")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    } else if (stype === "memRequested") {
      SearchType = "memRequested";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/usage/data/reservation")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    } else if (stype === "cpuRequested") {
      SearchType = "cpuRequested";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/usage/data/reservation")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    } else if (stype === "memLimit") {
      SearchType = "memLimit";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/reservation")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    } else if (stype === "cpuLimit") {
      SearchType = "cpuLimit";

      d3.selectAll("svg").each(function () {
        d3.select(this).remove();
      });

      fetch("/reservation")
        .then((res) => res.json())
        .then((res) => {
          drawTreeMap(res);
        });
    }
  };
}

/*  Version vom ddd */
/*
d3.select("#SearchType").on("input", setSearchType);

var SearchType = "mem";

function setSearchType() {
  var _value = d3.select("#SearchType").property("value");
  //console.log("#SearchType "+ _value)
  if (_value === "Memory") {
    SearchType = "mem";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("http://localhost:8180/usage/memory")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  } else if (_value === "Cpu") {
    SearchType = "cpu";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("http://localhost:8180/usage/memory")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  } else if (_value === "memRequested") {
    SearchType = "memRequested";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("http://localhost:8180/reservation")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  } else if (_value === "cpuRequested") {
    SearchType = "cpuRequested";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("http://localhost:8180/reservation")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  } else if (_value === "memLimit") {
    SearchType = "memLimit";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("/reservation")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  } else if (_value === "cpuLimit") {
    SearchType = "cpuLimit";

    d3.selectAll("svg").each(function(){
      d3.select(this).remove();
    });

    fetch("/reservation")
      .then((res) => res.json())
      .then((res) => {
        drawTreeMap(res);
      });
  }
}
*/

// Größe des Chart Fenster DIv an die Fenstergröße anpassen

var viewWidth = window.innerWidth;
var viewHeight = window.innerHeight;

console.log("viewWidth: " + viewWidth);
console.log("viewHeight: " + viewHeight);

if (viewHeight <= 800) {
  d3.select("#d3usageTreeChart").style("height", "400px");
} else {
  d3.select("#d3usageTreeChart").style("height", "600px");
}

// Größe des Charts an die größe des des divs anpassen
// div größe auslesen
var chartdiv = d3.select("#d3usageTreeChart");
var chartwidth = chartdiv.style("width").replace("px", "");
var chartheight = chartdiv.style("height").replace("px", "");

// debug logging
console.log("div Breite: " + chartwidth);
console.log("div Höhe: " + chartheight);

const margin = { top: 20, bottom: 20, left: 20, right: 20 };
//const width = 1600 - margin.left - margin.right;
const width = chartwidth - margin.left - margin.right ;
//const height = 800 - margin.top - margin.bottom;
const height = chartheight - margin.top - margin.bottom;

console.log("chart Breite: " + width);
console.log("chart Höhe: " + height);

const drawTreeMap = (dataset) => {
  console.log("Searchtype = " + SearchType);
  const hierarchy = d3
      .hierarchy(dataset)
      .sum(function (d) {
        //console.log("d: "+d)
        //console.dir(d)
        if (SearchType === "mem") {
          return d.mem;
        } else if (SearchType === "cpu") {
          return d.cpu;
        } else if (SearchType == "memRequested") {
          return d.memRequested;
        } else if (SearchType == "cpuRequested") {
          return d.cpuRequested;
        } else if (SearchType == "memLimit") {
          return d.memLimit;
        } else if (SearchType == "cpuLimit") {
          return d.cpuLimit;
        }
      })
      .sort(function (a, b) {
        //console.log("b: "+b + " ; "+ "a: "  + a)
        //console.dir(b)
        //console.dir(a)
        if (SearchType === "mem") {
          return b.mem - a.mem;
        } else if (SearchType === "cpu") {
          return b.cpu - a.cpu;
        } else if (SearchType === "memRequested") {
          return b.memRequested - a.memRequested;
        } else if (SearchType === "cpuRequested") {
          return b.cpuRequested - a.cpuRequested;
        } else if (SearchType === "memLimit") {
          return b.memLimit - a.memLimit;
        } else if (SearchType === "cpuLimit") {
          return b.cpuLimit - a.cpuLimit;
        }
      }),
    treemap = d3
      .treemap()
      .size([width, height])
      .paddingTop(2)
      .paddingRight(1)
      .paddingInner(0)
      .round(true);

  root = treemap(hierarchy);

  const categories = dataset.children.map((d) => d.name);

  const colors = [
    "#51574a",
    "#447c69",
    "#74c493",
    "#8e8c6d",
    "#e4bf80",
    "#e9d78e",
    "#e2975d",
    "#f19670",
    "#e16552",
    "#c94a53",
    "#be5168",
    "#a34974",
    "#993767",
    "#65387d",
    "#4e2472",
    "#9163b6",
    "#e279a3",
    "#e0598b",
    "#7c9fb0",
    "#5698c4",
    "#9abf88",
  ];
  /*
  const colors = [
    "#8dd3c7",    "#ffffb3",    "#bebada",
    "#fb8072",    "#80b1d3",    "#fdb462",
    "#b3de69",    "#fccde5",    "#d9d9d9",
    "#bc80bd",    "#ccebc5",    "#ffed6f",
    "#66c2a5",    "#fc8d62",    "#8da0cb",
    "#e78ac3",    "#a6d854",    "#ffd92f",
    "#e5c494",    "#b3b3b3",
  ];
*/
  colorScale = d3
    .scaleOrdinal() // the scale function
    .domain(categories) // the data
    .range(colors); // the way the data should be shown

  var opacity = d3.scaleLinear().domain([10, 30]).range([0.8, 1]);

  const svg = d3
    .select("#d3usageTreeChart")
    .append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
    .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

  svg
    .selectAll("rect")
    .data(root.leaves())
    .enter()
    .append("rect")
    .attr("x", (d) => d.x0)
    .attr("y", (d) => d.y0)
    .attr("width", (d) => d.x1 - d.x0)
    .attr("height", (d) => d.y1 - d.y0)
    .attr("fill", (d) => colorScale(d.data.name))
    .attr("fill", function (d) {
      return colorScale(d.parent.parent.data.name);
    })
    /*
    .attr("opacity", function (d) {
      return opacity(d.data.mem);
    })
    */
    .on("mousemove", showtooltip)
    .on("mouseout", hidetooltip);

  /* 
  // and to add the text labels
  svg
    .selectAll("text")
    .data(root.leaves())
    .enter()
    .append("text")
    .attr("x", function (d) {
      return d.x0 + 5;
    }) // +10 to adjust position (more right)
    .attr("y", function (d) {
      return d.y0 + 10;
    }) // +20 to adjust position (lower)
    .text(function (d) {
      return "p : " + d.parent.data.name;
    })
    .attr("font-size", "6px")
    .attr("fill", "black")
    .append("svg:tspan")
    .text(function (d) {
      return "c : " + d.data.name;
    })
    .attr("x", function (d) {
      return d.x0 + 5;
    })
    .attr("y", function (d) {
      return d.y0 + 16;
    })
    .attr("font-size", "6px")
    .attr("fill", "black");
*/

  // and to add the text labels for Values
  /*
  svg
    .selectAll("vals")
    .data(root.leaves())
    .enter()
    .append("text")
    .attr("x", function (d) {
      return d.x0 + 5;
    }) // +10 to adjust position (more right)
    .attr("y", function (d) {
      return d.y0 + 23;
    }) // +20 to adjust position (lower)
    .text(function (d) {
      if (SearchType === "mem") {
        return d.data.mem + "Ki";
      } else if (SearchType === "cpu") {
        return d.data.cpu + " mCpu";
      }
    })
    .attr("font-size", "6px")
    .attr("fill", "black");
*/

  /*

  // Add title for the 3 groups
  svg
    .selectAll("titles")
    .data(
      root.descendants().filter(function (d) {
        return d.depth == 1;
      })
    )
    .enter()
    .append("text")
    .attr("x", function (d) {
      return d.x0;
    })
    .attr("y", function (d) {
      return d.y0 + 21;
    })
    .text(function (d) {
      return d.data.name;
    })
    .attr("font-size", "10px");
  // .attr("fill",  function(d){ return colorScale(d.data.name)} );
*/

  // Add title

  d3.select("#d3usgaeTreeWindow > div.card-header.py-3.d-flex.flex-row.align-items-center.justify-content-between > h6")
    .text(function () {
      if (SearchType === "mem") {
        return "Memory Consumption per pod";
      } else if (SearchType === "cpu") {
        return "Cpu Consumption per pod";
      } else if (SearchType === "memRequested") {
        return "Memory Reservation";
      } else if (SearchType === "cpuRequested") {
        return "Cpu Reservation";
      } else if (SearchType === "memLimit") {
        return "Memory Limit";
      } else if (SearchType === "cpuLimit") {
        return "Cpu Limit";
      } else {
        return "Unknown Serach Type " + SearchType;
      }
    })
  /*
  svg
    .append("text")
    .attr("x", 0)
    .attr("y", 0) // +20 to adjust position (lower)
    .text(function () {
      if (SearchType === "mem") {
        return "Memory Consumption per pod";
      } else if (SearchType === "cpu") {
        return "Cpu Consumption per pod";
      } else if (SearchType === "memRequested") {
        return "Memory Reservation";
      } else if (SearchType === "cpuRequested") {
        return "Cpu Reservation";
      } else {
        return "Unknown Serach Type " + SearchType;
      }
    })

    .attr("font-size", "19px")
    .attr("fill", "grey");

    */


  // Add legend

  const legendPadding = {
    left: 10,
    top: 10,
  };
  const legendRectSizes = {
    width: 20,
    height: 20,
  };

  const leg = d3
    .select("#legende")
    .append("svg")
    .attr("width", width + margin.left + margin.right)
    .style("padding-left", "20px")
    .attr("height", 250)
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");
  //.attr("height", height + margin.top + margin.bottom);
  //.append("g")
  //.attr("transform", "translate(" + 30 + "," + 810 + ")");

  const x = leg.append("g").attr("id", "legend");

  //const x = svg.append("g").attr("id", "legend");
  x.append("rect")
    .attr("y", 0)
    // .attr("y", height + legendRectSizes.height)
    .attr("width", width)
    .attr("height", legendRectSizes.height * 2)
    .attr("fill", bodyColor);

  x.append("g")
    .selectAll("rect")
    .attr("id", "legend")
    .data(root.children)
    .enter()
    .append("rect")
    .attr("class", "legend-item")
    .attr("y", legendPadding.top)
    // .attr("y", height + legendRectSizes.height)
    .attr("x", (d, i) => (margin.left + 10) * i)
    .attr("width", legendRectSizes.width)
    .attr("height", legendRectSizes.height)
    // .attr("fill", (d) => colorScale(d.children[0].value));
    .attr("fill", (d) => colorScale(d.data.name));

  x.append("g")
    .selectAll("text")
    .data(root.children)
    .enter()
    .append("text")
    .attr("class", "legend-text")
    .attr("font-size", "12px")
    .attr("fill", "#465353")
    //.attr("y", 5 )
    //.attr("y", height + legendRectSizes.height * 1.5)
    .attr("y", (d, i) => ( margin.left + 10) * i * -1 - legendRectSizes.width)
    .attr("x", legendRectSizes.height + legendPadding.top +5)
    //.attr("x", (d, i) => margin.left * i + legendRectSizes.width)
    .attr("text-achor", "start")
    .attr("transform", "translate(-10,0) rotate(90)")
    //.attr("transform", "rotate(90)")
    .text((d) => d.data.name);
};

// show tooltip

function showtooltip(event, d) {
  // tooltip position in abhängigkeit der sidebar breite
  /*
  var sideBardiv = d3.select("#accordionSidebar");
  var sideBarWidth = sideBardiv.style("width").replace("px", "");
  console.log("sideBarWidth: " + sideBarWidth);

  */

  // tooltip position in abhängigkeit der position des Charts
  var div = d3.select("#d3usageTreeChart");
  var chartpos = div.node().getBoundingClientRect();

  /*
  console.log("Oben: " + chartpos.top + "px");
  console.log("Links: " + chartpos.left + "px");
*/

  var coords = d3.pointer(event);
  var xPos = coords[0] + chartpos.left - 100;
  var yPos = coords[1] + chartpos.top + 50;

  divToolTip.transition().duration(200).style("opacity", 0.95);
  divToolTip
    .style("left", xPos + "px")
    .style("top", yPos + "px")
    .style("font-size", "12px")
    .style("width", "250px")
    .style("height", "150px")
    .html(function () {
      if (SearchType === "mem" || SearchType === "cpu") {
        return (
          "<strong>Pod: " +
          d.parent.data.name +
          "</strong><br/>" +
          "Container: " +
          d.data.name +
          "<br/>" +
          "Namespace: " +
          d.parent.parent.data.name +
          "<br/>" +
          "Memory: " +
          Math.round(d.data.mem / 1024) +
          " MiByte<br/>" +
          "CPU: " +
          Math.round(d.data.cpu / 1000 / 1000) +
          " mCpu"
        );
      } else {
        return (
          "<strong>Pod: " +
          d.parent.data.name +
          "</strong><br/>" +
          "Container: " +
          d.data.name +
          "<br/>" +
          "Namespace: " +
          d.parent.parent.data.name +
          "<br/>" +
          "MemRes: " +
          Math.round(d.data.memRequested / 1024) +
          "MiByte<br/>" +
          "MemLimit: " +
          Math.round(d.data.memLimit / 1024) +
          "MiByte<br/>" +
          "CpuRes: " +
          Math.round(d.data.cpuRequested) +
          "mCpu<br/>" +
          "CpuLimit: " +
          Math.round(d.data.cpuLimit) +
          "mCpu"
        );
      }
    });
}
// hide tooltip
function hidetooltip(d) {
  divToolTip.transition().duration(500).style("opacity", 0);
}
function showErrorMsg(msg) {
  d3.select("#errorMsg").attr("style", "display: block");
  d3.select(".alert").html("<strong>Error:</strong>" + msg);
}
function hideErrorMsg() {
  d3.select("#errorMsg").attr("style", "display: none");
}
