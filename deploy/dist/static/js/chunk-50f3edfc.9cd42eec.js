(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-50f3edfc"],{c361:function(e,t,l){"use strict";l.r(t);var a=function(){var e=this,t=e.$createElement,l=e._self._c||t;return l("div",{staticClass:"components-container"},[l("el-drag-select",{staticStyle:{width:"500px"},attrs:{multiple:"",placeholder:"请选择"},model:{value:e.value,callback:function(t){e.value=t},expression:"value"}},e._l(e.options,(function(e){return l("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1),l("div",{staticStyle:{"margin-top":"30px"}},e._l(e.value,(function(t){return l("el-tag",{key:t,staticStyle:{"margin-right":"15px"}},[e._v(" "+e._s(t)+" ")])})),1)],1)},n=[],s=function(){var e=this,t=e.$createElement,l=e._self._c||t;return l("el-select",e._g(e._b({ref:"dragSelect",staticClass:"drag-select",attrs:{multiple:""},model:{value:e.selectVal,callback:function(t){e.selectVal=t},expression:"selectVal"}},"el-select",e.$attrs,!1),e.$listeners),[e._t("default")],2)},r=[],c=(l("a434"),l("2909")),u=l("8a23"),i={name:"DragSelect",props:{value:{type:Array,required:!0}},computed:{selectVal:{get:function(){return Object(c["a"])(this.value)},set:function(e){this.$emit("input",Object(c["a"])(e))}}},mounted:function(){this.setSort()},methods:{setSort:function(){var e=this,t=this.$refs.dragSelect.$el.querySelectorAll(".el-select__tags > span")[0];this.sortable=u["a"].create(t,{ghostClass:"sortable-ghost",setData:function(e){e.setData("Text","")},onEnd:function(t){var l=e.value.splice(t.oldIndex,1)[0];e.value.splice(t.newIndex,0,l)}})}}},o=i,p=(l("c482"),l("2877")),v=Object(p["a"])(o,s,r,!1,null,"7e74c641",null),f=v.exports,d={name:"DragSelectDemo",components:{ElDragSelect:f},data:function(){return{value:["Apple","Banana","Orange"],options:[{value:"Apple",label:"Apple"},{value:"Banana",label:"Banana"},{value:"Orange",label:"Orange"},{value:"Pear",label:"Pear"},{value:"Strawberry",label:"Strawberry"}]}}},b=d,g=Object(p["a"])(b,a,n,!1,null,null,null);t["default"]=g.exports},c482:function(e,t,l){"use strict";l("f302")},f302:function(e,t,l){}}]);