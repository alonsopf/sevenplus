/**
 * central-click.js
 */
 function cargaDistritos() {
  $("#mainContent").html('<canvas height="1300" width="1500" style="-webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas>');
  var canvas = document.querySelector("canvas"),
                width = 1500,
                height = 1300;
              var o = [];
              var i = [];
               o.push({
                    id: 0,
                    name: "Asociacion Baja Norte",
                    budget : "",
                    DT : "AGM411"
                });
              $.ajax({
        url: "/listaDistritosPresupuesto",
        dataType : "json",
        type : "post",
        async : true,
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp) {
          if(resp.success==1)
          {
var ii = 0, aver2 = Object.keys(resp.distritos);
var con = 1;
            for(ii=0;ii<aver2.length;ii++)
            {
              o.push({
                    id: con,
                    name: resp.distritos[aver2[ii]].Name.trim().replace("DS-",""),
                    budget : "$200,000.00",
                    DT : aver2[ii]
                });
               i.push({
                    source: 0,
                    target: con
                });
               
              con++;
            }  

                initColor(o), initRadius(o);
                var s = new relation;
                s.setNodes(o), s.setLinks(i), s.setCanvas(canvas), s.setSize(width, height), s.setRadius(12), s.setLinkLength(90), s.setCharge(-60), s.init(), s.run();
         
          }
        }
    });
     
            
 }
d4.svg.BubbleChart.define("central-click", function (options) {
  var self = this;
  self.setup = (function (node) {
    var original = self.setup;
    return function (node) {
      var fn = original.apply(this, arguments);
      self.event.on("click", function(node) {
        if (node.selectAll("text.central-click")[0].length === 1) {
          console.log(node);
          cargaDistritos();
          //alert("Hello there!\nCentral bubble is clicked.");
        }
      });
      return fn;
    };
  })();

  self.reset = (function (node) {
    var original = self.reset;
    return function (node) {
      var fn = original.apply(this, arguments);
      node.select("text.central-click").remove();
      return fn;
    };
  })();

  self.moveToCentral = (function (node) {
    var original = self.moveToCentral;
    return function (node) {
      var fn = original.apply(this, arguments);
      var transition = self.getTransition().centralNode;
      var i;
      transition.each("end", function() {
        node.append("text").classed({"central-click": true})
          .attr(options.attr)
          .style(options.style)
          .attr("x", function (d) {return d.cx;})
          .attr("y", function (d) {return d.cy;})
          .text(options.text)
          .style("opacity", 0).transition().duration(self.getOptions().transitDuration / 2).style("opacity", "0.8");
      });
      
      return fn;
    };
  })();
});
var transitionGlobal;
/**
 * Text lines
 *
 * @module
 *
 * @param {object} options - setting of text lines
 * @param {object[]} options.format - n-th object is used to format n-th line
 * @param {string} options.format.textField - name of property will be used in function text()
 * @param {string} options.format.classed - classes used to apply to [text](http://www.w3.org/TR/SVG/text.html#TextElement)
 * @param {string} options.format.style - style used to apply to [text](http://www.w3.org/TR/SVG/text.html#TextElement)
 * @param {string} options.format.attr - attribute used to apply to [text](http://www.w3.org/TR/SVG/text.html#TextElement)
 * @param {object[]} options.centralFormat - like #format but used to format central-bubble
 */
d4.svg.BubbleChart.define("lines", function (options) {
  /*
   * @param
   *  options = {
   *    format: [ //n-th object is used to format n-th line
   *      {
   *        textField: #string, name of property will be used in function text(), @link
   *        classed: #object, @link
   *        style: #object, @link
   *        attr: #object, @link
   *      }
   *    ],
   *    centralFormat: [ //@see #format, but used to format central-bubble
   *    ]
   *  }
   * */

  var self = this;

  self.setup = (function () {
    var original = self.setup;
    return function () {
      var fn = original.apply(this, arguments);
      var node = self.getNodes();
      var i;
     /* for(i=0;i<options.format.length;i++){
        var f = options.format[i];
        node.append("text")
          .classed(f.classed)
          .style(f.style)
          .attr(f.attr)
          .text(function (d) {return d.item[f.textField];});
      }*/
      
      $.each(options.format, function (i, f) {
        node.append("text")
          .classed(f.classed)
          .style(f.style)
          .attr(f.attr)
          .text(function (d) {return d.item[f.textField];});
      });
      
      return fn;
    };
  })();

  self.reset = (function (node) {
    var original = self.reset;
    return function (node) {
      var fn = original.apply(this, arguments);
     /* for(i=0;i<options.format.length;i++){
        var f = options.format[i];
        var tNode = d4.select(node.selectAll("text")[0][i]);
        tNode.classed(f.classed).text(function (d) {return d.item[f.textField];})
          .transition().duration(self.getOptions().transitDuration)
          .style(f.style)
          .attr(f.attr);
      }*/
      
      $.each(options.format, function (i, f) {
        var tNode = d4.select(node.selectAll("text")[0][i]);
        tNode.classed(f.classed).text(function (d) {return d.item[f.textField];})
          .transition().duration(self.getOptions().transitDuration)
          .style(f.style)
          .attr(f.attr);
      });
      
      return fn;
    };
  })();

  self.moveToCentral = (function (node) {
    var original = self.moveToCentral;
    return function (node) {
      var fn = original.apply(this, arguments);
    /*  for(i=0;i<options.format.length;i++){
        var f = options.format[i];
          var tNode = d4.select(node.selectAll("text")[0][i]);
        tNode.transition().duration(self.getOptions().transitDuration)
          .style(f.style)
          .attr(f.attr);
        f.classed !== undefined && tNode.classed(f.classed);
        f.textField !== undefined && tNode.text(function (d) {return d.item[f.textField];});
      }*/
      
      $.each(options.centralFormat, function (i, f) {
        var tNode = d4.select(node.selectAll("text")[0][i]);
        tNode.transition().duration(self.getOptions().transitDuration)
          .style(f.style)
          .attr(f.attr);
        f.classed !== undefined && tNode.classed(f.classed);
        f.textField !== undefined && tNode.text(function (d) {return d.item[f.textField];});
      });
      
      return fn;
    };
  })();
});