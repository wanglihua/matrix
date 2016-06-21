
/*!
 DataTables 1.10.12
 ©2008-2015 SpryMedia Ltd - datatables.net/license
*/
(function(h){"function"===typeof define&&define.amd?define(["jquery"],function(D){return h(D,window,document)}):"object"===typeof exports?module.exports=function(D,I){D||(D=window);I||(I="undefined"!==typeof window?require("jquery"):require("jquery")(D));return h(I,D,D.document)}:h(jQuery,window,document)})(function(h,D,I,k){function X(a){var b,c,d={};h.each(a,function(e){if((b=e.match(/^([^A-Z]+?)([A-Z])/))&&-1!=="a aa ai ao as b fn i m o s ".indexOf(b[1]+" "))c=e.replace(b[0],b[2].toLowerCase()),
d[c]=e,"o"===b[1]&&X(a[e])});a._hungarianMap=d}function K(a,b,c){a._hungarianMap||X(a);var d;h.each(b,function(e){d=a._hungarianMap[e];if(d!==k&&(c||b[d]===k))"o"===d.charAt(0)?(b[d]||(b[d]={}),h.extend(!0,b[d],b[e]),K(a[d],b[d],c)):b[d]=b[e]})}function Da(a){var b=m.defaults.oLanguage,c=a.sZeroRecords;!a.sEmptyTable&&(c&&"No data available in table"===b.sEmptyTable)&&E(a,a,"sZeroRecords","sEmptyTable");!a.sLoadingRecords&&(c&&"Loading..."===b.sLoadingRecords)&&E(a,a,"sZeroRecords","sLoadingRecords");
a.sInfoThousands&&(a.sThousands=a.sInfoThousands);(a=a.sDecimal)&&db(a)}function eb(a){A(a,"ordering","bSort");A(a,"orderMulti","bSortMulti");A(a,"orderClasses","bSortClasses");A(a,"orderCellsTop","bSortCellsTop");A(a,"order","aaSorting");A(a,"orderFixed","aaSortingFixed");A(a,"paging","bPaginate");A(a,"pagingType","sPaginationType");A(a,"pageLength","iDisplayLength");A(a,"searching","bFilter");"boolean"===typeof a.sScrollX&&(a.sScrollX=a.sScrollX?"100%":"");"boolean"===typeof a.scrollX&&(a.scrollX=
a.scrollX?"100%":"");if(a=a.aoSearchCols)for(var b=0,c=a.length;b<c;b++)a[b]&&K(m.models.oSearch,a[b])}function fb(a){A(a,"orderable","bSortable");A(a,"orderData","aDataSort");A(a,"orderSequence","asSorting");A(a,"orderDataType","sortDataType");var b=a.aDataSort;b&&!h.isArray(b)&&(a.aDataSort=[b])}function gb(a){if(!m.__browser){var b={};m.__browser=b;var c=h("<div/>").css({position:"fixed",top:0,left:0,height:1,width:1,overflow:"hidden"}).append(h("<div/>").css({position:"absolute",top:1,left:1,
width:100,overflow:"scroll"}).append(h("<div/>").css({width:"100%",height:10}))).appendTo("body"),d=c.children(),e=d.children();b.barWidth=d[0].offsetWidth-d[0].clientWidth;b.bScrollOversize=100===e[0].offsetWidth&&100!==d[0].clientWidth;b.bScrollbarLeft=1!==Math.round(e.offset().left);b.bBounding=c[0].getBoundingClientRect().width?!0:!1;c.remove()}h.extend(a.oBrowser,m.__browser);a.oScroll.iBarWidth=m.__browser.barWidth}function hb(a,b,c,d,e,f){var g,j=!1;c!==k&&(g=c,j=!0);for(;d!==e;)a.hasOwnProperty(d)&&
(g=j?b(g,a[d],d,a):a[d],j=!0,d+=f);return g}function Ea(a,b){var c=m.defaults.column,d=a.aoColumns.length,c=h.extend({},m.models.oColumn,c,{nTh:b?b:I.createElement("th"),sTitle:c.sTitle?c.sTitle:b?b.innerHTML:"",aDataSort:c.aDataSort?c.aDataSort:[d],mData:c.mData?c.mData:d,idx:d});a.aoColumns.push(c);c=a.aoPreSearchCols;c[d]=h.extend({},m.models.oSearch,c[d]);ja(a,d,h(b).data())}function ja(a,b,c){var b=a.aoColumns[b],d=a.oClasses,e=h(b.nTh);if(!b.sWidthOrig){b.sWidthOrig=e.attr("width")||null;var f=
(e.attr("style")||"").match(/width:\s*(\d+[pxem%]+)/);f&&(b.sWidthOrig=f[1])}c!==k&&null!==c&&(fb(c),K(m.defaults.column,c),c.mDataProp!==k&&!c.mData&&(c.mData=c.mDataProp),c.sType&&(b._sManualType=c.sType),c.className&&!c.sClass&&(c.sClass=c.className),h.extend(b,c),E(b,c,"sWidth","sWidthOrig"),c.iDataSort!==k&&(b.aDataSort=[c.iDataSort]),E(b,c,"aDataSort"));var g=b.mData,j=Q(g),i=b.mRender?Q(b.mRender):null,c=function(a){return"string"===typeof a&&-1!==a.indexOf("@")};b._bAttrSrc=h.isPlainObject(g)&&
(c(g.sort)||c(g.type)||c(g.filter));b._setter=null;b.fnGetData=function(a,b,c){var d=j(a,b,k,c);return i&&b?i(d,b,a,c):d};b.fnSetData=function(a,b,c){return R(g)(a,b,c)};"number"!==typeof g&&(a._rowReadObject=!0);a.oFeatures.bSort||(b.bSortable=!1,e.addClass(d.sSortableNone));a=-1!==h.inArray("asc",b.asSorting);c=-1!==h.inArray("desc",b.asSorting);!b.bSortable||!a&&!c?(b.sSortingClass=d.sSortableNone,b.sSortingClassJUI=""):a&&!c?(b.sSortingClass=d.sSortableAsc,b.sSortingClassJUI=d.sSortJUIAscAllowed):
!a&&c?(b.sSortingClass=d.sSortableDesc,b.sSortingClassJUI=d.sSortJUIDescAllowed):(b.sSortingClass=d.sSortable,b.sSortingClassJUI=d.sSortJUI)}function Y(a){if(!1!==a.oFeatures.bAutoWidth){var b=a.aoColumns;Fa(a);for(var c=0,d=b.length;c<d;c++)b[c].nTh.style.width=b[c].sWidth}b=a.oScroll;(""!==b.sY||""!==b.sX)&&ka(a);u(a,null,"column-sizing",[a])}function Z(a,b){var c=la(a,"bVisible");return"number"===typeof c[b]?c[b]:null}function $(a,b){var c=la(a,"bVisible"),c=h.inArray(b,c);return-1!==c?c:null}
function aa(a){var b=0;h.each(a.aoColumns,function(a,d){d.bVisible&&"none"!==h(d.nTh).css("display")&&b++});return b}function la(a,b){var c=[];h.map(a.aoColumns,function(a,e){a[b]&&c.push(e)});return c}function Ga(a){var b=a.aoColumns,c=a.aoData,d=m.ext.type.detect,e,f,g,j,i,h,l,q,t;e=0;for(f=b.length;e<f;e++)if(l=b[e],t=[],!l.sType&&l._sManualType)l.sType=l._sManualType;else if(!l.sType){g=0;for(j=d.length;g<j;g++){i=0;for(h=c.length;i<h;i++){t[i]===k&&(t[i]=B(a,i,e,"type"));q=d[g](t[i],a);if(!q&&
g!==d.length-1)break;if("html"===q)break}if(q){l.sType=q;break}}l.sType||(l.sType="string")}}function ib(a,b,c,d){var e,f,g,j,i,n,l=a.aoColumns;if(b)for(e=b.length-1;0<=e;e--){n=b[e];var q=n.targets!==k?n.targets:n.aTargets;h.isArray(q)||(q=[q]);f=0;for(g=q.length;f<g;f++)if("number"===typeof q[f]&&0<=q[f]){for(;l.length<=q[f];)Ea(a);d(q[f],n)}else if("number"===typeof q[f]&&0>q[f])d(l.length+q[f],n);else if("string"===typeof q[f]){j=0;for(i=l.length;j<i;j++)("_all"==q[f]||h(l[j].nTh).hasClass(q[f]))&&
d(j,n)}}if(c){e=0;for(a=c.length;e<a;e++)d(e,c[e])}}function N(a,b,c,d){var e=a.aoData.length,f=h.extend(!0,{},m.models.oRow,{src:c?"dom":"data",idx:e});f._aData=b;a.aoData.push(f);for(var g=a.aoColumns,j=0,i=g.length;j<i;j++)g[j].sType=null;a.aiDisplayMaster.push(e);b=a.rowIdFn(b);b!==k&&(a.aIds[b]=f);(c||!a.oFeatures.bDeferRender)&&Ha(a,e,c,d);return e}function ma(a,b){var c;b instanceof h||(b=h(b));return b.map(function(b,e){c=Ia(a,e);return N(a,c.data,e,c.cells)})}function B(a,b,c,d){var e=a.iDraw,
f=a.aoColumns[c],g=a.aoData[b]._aData,j=f.sDefaultContent,i=f.fnGetData(g,d,{settings:a,row:b,col:c});if(i===k)return a.iDrawError!=e&&null===j&&(L(a,0,"Requested unknown parameter "+("function"==typeof f.mData?"{function}":"'"+f.mData+"'")+" for row "+b+", column "+c,4),a.iDrawError=e),j;if((i===g||null===i)&&null!==j&&d!==k)i=j;else if("function"===typeof i)return i.call(g);return null===i&&"display"==d?"":i}function jb(a,b,c,d){a.aoColumns[c].fnSetData(a.aoData[b]._aData,d,{settings:a,row:b,col:c})}
function Ja(a){return h.map(a.match(/(\\.|[^\.])+/g)||[""],function(a){return a.replace(/\\./g,".")})}function Q(a){if(h.isPlainObject(a)){var b={};h.each(a,function(a,c){c&&(b[a]=Q(c))});return function(a,c,f,g){var j=b[c]||b._;return j!==k?j(a,c,f,g):a}}if(null===a)return function(a){return a};if("function"===typeof a)return function(b,c,f,g){return a(b,c,f,g)};if("string"===typeof a&&(-1!==a.indexOf(".")||-1!==a.indexOf("[")||-1!==a.indexOf("("))){var c=function(a,b,f){var g,j;if(""!==f){j=Ja(f);
for(var i=0,n=j.length;i<n;i++){f=j[i].match(ba);g=j[i].match(U);if(f){j[i]=j[i].replace(ba,"");""!==j[i]&&(a=a[j[i]]);g=[];j.splice(0,i+1);j=j.join(".");if(h.isArray(a)){i=0;for(n=a.length;i<n;i++)g.push(c(a[i],b,j))}a=f[0].substring(1,f[0].length-1);a=""===a?g:g.join(a);break}else if(g){j[i]=j[i].replace(U,"");a=a[j[i]]();continue}if(null===a||a[j[i]]===k)return k;a=a[j[i]]}}return a};return function(b,e){return c(b,e,a)}}return function(b){return b[a]}}function R(a){if(h.isPlainObject(a))return R(a._);
if(null===a)return function(){};if("function"===typeof a)return function(b,d,e){a(b,"set",d,e)};if("string"===typeof a&&(-1!==a.indexOf(".")||-1!==a.indexOf("[")||-1!==a.indexOf("("))){var b=function(a,d,e){var e=Ja(e),f;f=e[e.length-1];for(var g,j,i=0,n=e.length-1;i<n;i++){g=e[i].match(ba);j=e[i].match(U);if(g){e[i]=e[i].replace(ba,"");a[e[i]]=[];f=e.slice();f.splice(0,i+1);g=f.join(".");if(h.isArray(d)){j=0;for(n=d.length;j<n;j++)f={},b(f,d[j],g),a[e[i]].push(f)}else a[e[i]]=d;return}j&&(e[i]=e[i].replace(U,
""),a=a[e[i]](d));if(null===a[e[i]]||a[e[i]]===k)a[e[i]]={};a=a[e[i]]}if(f.match(U))a[f.replace(U,"")](d);else a[f.replace(ba,"")]=d};return function(c,d){return b(c,d,a)}}return function(b,d){b[a]=d}}function Ka(a){return G(a.aoData,"_aData")}function na(a){a.aoData.length=0;a.aiDisplayMaster.length=0;a.aiDisplay.length=0;a.aIds={}}function oa(a,b,c){for(var d=-1,e=0,f=a.length;e<f;e++)a[e]==b?d=e:a[e]>b&&a[e]--; -1!=d&&c===k&&a.splice(d,1)}function ca(a,b,c,d){var e=a.aoData[b],f,g=function(c,d){for(;c.childNodes.length;)c.removeChild(c.firstChild);
c.innerHTML=B(a,b,d,"display")};if("dom"===c||(!c||"auto"===c)&&"dom"===e.src)e._aData=Ia(a,e,d,d===k?k:e._aData).data;else{var j=e.anCells;if(j)if(d!==k)g(j[d],d);else{c=0;for(f=j.length;c<f;c++)g(j[c],c)}}e._aSortData=null;e._aFilterData=null;g=a.aoColumns;if(d!==k)g[d].sType=null;else{c=0;for(f=g.length;c<f;c++)g[c].sType=null;La(a,e)}}function Ia(a,b,c,d){var e=[],f=b.firstChild,g,j,i=0,n,l=a.aoColumns,q=a._rowReadObject,d=d!==k?d:q?{}:[],t=function(a,b){if("string"===typeof a){var c=a.indexOf("@");
-1!==c&&(c=a.substring(c+1),R(a)(d,b.getAttribute(c)))}},S=function(a){if(c===k||c===i)j=l[i],n=h.trim(a.innerHTML),j&&j._bAttrSrc?(R(j.mData._)(d,n),t(j.mData.sort,a),t(j.mData.type,a),t(j.mData.filter,a)):q?(j._setter||(j._setter=R(j.mData)),j._setter(d,n)):d[i]=n;i++};if(f)for(;f;){g=f.nodeName.toUpperCase();if("TD"==g||"TH"==g)S(f),e.push(f);f=f.nextSibling}else{e=b.anCells;f=0;for(g=e.length;f<g;f++)S(e[f])}if(b=b.firstChild?b:b.nTr)(b=b.getAttribute("id"))&&R(a.rowId)(d,b);return{data:d,cells:e}}
function Ha(a,b,c,d){var e=a.aoData[b],f=e._aData,g=[],j,i,n,l,q;if(null===e.nTr){j=c||I.createElement("tr");e.nTr=j;e.anCells=g;j._DT_RowIndex=b;La(a,e);l=0;for(q=a.aoColumns.length;l<q;l++){n=a.aoColumns[l];i=c?d[l]:I.createElement(n.sCellType);i._DT_CellIndex={row:b,column:l};g.push(i);if((!c||n.mRender||n.mData!==l)&&(!h.isPlainObject(n.mData)||n.mData._!==l+".display"))i.innerHTML=B(a,b,l,"display");n.sClass&&(i.className+=" "+n.sClass);n.bVisible&&!c?j.appendChild(i):!n.bVisible&&c&&i.parentNode.removeChild(i);
n.fnCreatedCell&&n.fnCreatedCell.call(a.oInstance,i,B(a,b,l),f,b,l)}u(a,"aoRowCreatedCallback",null,[j,f,b])}e.nTr.setAttribute("role","row")}function La(a,b){var c=b.nTr,d=b._aData;if(c){var e=a.rowIdFn(d);e&&(c.id=e);d.DT_RowClass&&(e=d.DT_RowClass.split(" "),b.__rowc=b.__rowc?pa(b.__rowc.concat(e)):e,h(c).removeClass(b.__rowc.join(" ")).addClass(d.DT_RowClass));d.DT_RowAttr&&h(c).attr(d.DT_RowAttr);d.DT_RowData&&h(c).data(d.DT_RowData)}}function kb(a){var b,c,d,e,f,g=a.nTHead,j=a.nTFoot,i=0===
h("th, td",g).length,n=a.oClasses,l=a.aoColumns;i&&(e=h("<tr/>").appendTo(g));b=0;for(c=l.length;b<c;b++)f=l[b],d=h(f.nTh).addClass(f.sClass),i&&d.appendTo(e),a.oFeatures.bSort&&(d.addClass(f.sSortingClass),!1!==f.bSortable&&(d.attr("tabindex",a.iTabIndex).attr("aria-controls",a.sTableId),Ma(a,f.nTh,b))),f.sTitle!=d[0].innerHTML&&d.html(f.sTitle),Na(a,"header")(a,d,f,n);i&&da(a.aoHeader,g);h(g).find(">tr").attr("role","row");h(g).find(">tr>th, >tr>td").addClass(n.sHeaderTH);h(j).find(">tr>th, >tr>td").addClass(n.sFooterTH);
if(null!==j){a=a.aoFooter[0];b=0;for(c=a.length;b<c;b++)f=l[b],f.nTf=a[b].cell,f.sClass&&h(f.nTf).addClass(f.sClass)}}function ea(a,b,c){var d,e,f,g=[],j=[],i=a.aoColumns.length,n;if(b){c===k&&(c=!1);d=0;for(e=b.length;d<e;d++){g[d]=b[d].slice();g[d].nTr=b[d].nTr;for(f=i-1;0<=f;f--)!a.aoColumns[f].bVisible&&!c&&g[d].splice(f,1);j.push([])}d=0;for(e=g.length;d<e;d++){if(a=g[d].nTr)for(;f=a.firstChild;)a.removeChild(f);f=0;for(b=g[d].length;f<b;f++)if(n=i=1,j[d][f]===k){a.appendChild(g[d][f].cell);
for(j[d][f]=1;g[d+i]!==k&&g[d][f].cell==g[d+i][f].cell;)j[d+i][f]=1,i++;for(;g[d][f+n]!==k&&g[d][f].cell==g[d][f+n].cell;){for(c=0;c<i;c++)j[d+c][f+n]=1;n++}h(g[d][f].cell).attr("rowspan",i).attr("colspan",n)}}}}function O(a){var b=u(a,"aoPreDrawCallback","preDraw",[a]);if(-1!==h.inArray(!1,b))C(a,!1);else{var b=[],c=0,d=a.asStripeClasses,e=d.length,f=a.oLanguage,g=a.iInitDisplayStart,j="ssp"==y(a),i=a.aiDisplay;a.bDrawing=!0;g!==k&&-1!==g&&(a._iDisplayStart=j?g:g>=a.fnRecordsDisplay()?0:g,a.iInitDisplayStart=
-1);var g=a._iDisplayStart,n=a.fnDisplayEnd();if(a.bDeferLoading)a.bDeferLoading=!1,a.iDraw++,C(a,!1);else if(j){if(!a.bDestroying&&!lb(a))return}else a.iDraw++;if(0!==i.length){f=j?a.aoData.length:n;for(j=j?0:g;j<f;j++){var l=i[j],q=a.aoData[l];null===q.nTr&&Ha(a,l);l=q.nTr;if(0!==e){var t=d[c%e];q._sRowStripe!=t&&(h(l).removeClass(q._sRowStripe).addClass(t),q._sRowStripe=t)}u(a,"aoRowCallback",null,[l,q._aData,c,j]);b.push(l);c++}}else c=f.sZeroRecords,1==a.iDraw&&"ajax"==y(a)?c=f.sLoadingRecords:
f.sEmptyTable&&0===a.fnRecordsTotal()&&(c=f.sEmptyTable),b[0]=h("<tr/>",{"class":e?d[0]:""}).append(h("<td />",{valign:"top",colSpan:aa(a),"class":a.oClasses.sRowEmpty}).html(c))[0];u(a,"aoHeaderCallback","header",[h(a.nTHead).children("tr")[0],Ka(a),g,n,i]);u(a,"aoFooterCallback","footer",[h(a.nTFoot).children("tr")[0],Ka(a),g,n,i]);d=h(a.nTBody);d.children().detach();d.append(h(b));u(a,"aoDrawCallback","draw",[a]);a.bSorted=!1;a.bFiltered=!1;a.bDrawing=!1}}function T(a,b){var c=a.oFeatures,d=c.bFilter;
c.bSort&&mb(a);d?fa(a,a.oPreviousSearch):a.aiDisplay=a.aiDisplayMaster.slice();!0!==b&&(a._iDisplayStart=0);a._drawHold=b;O(a);a._drawHold=!1}function nb(a){var b=a.oClasses,c=h(a.nTable),c=h("<div/>").insertBefore(c),d=a.oFeatures,e=h("<div/>",{id:a.sTableId+"_wrapper","class":b.sWrapper+(a.nTFoot?"":" "+b.sNoFooter)});a.nHolding=c[0];a.nTableWrapper=e[0];a.nTableReinsertBefore=a.nTable.nextSibling;for(var f=a.sDom.split(""),g,j,i,n,l,q,t=0;t<f.length;t++){g=null;j=f[t];if("<"==j){i=h("<div/>")[0];
n=f[t+1];if("'"==n||'"'==n){l="";for(q=2;f[t+q]!=n;)l+=f[t+q],q++;"H"==l?l=b.sJUIHeader:"F"==l&&(l=b.sJUIFooter);-1!=l.indexOf(".")?(n=l.split("."),i.id=n[0].substr(1,n[0].length-1),i.className=n[1]):"#"==l.charAt(0)?i.id=l.substr(1,l.length-1):i.className=l;t+=q}e.append(i);e=h(i)}else if(">"==j)e=e.parent();else if("l"==j&&d.bPaginate&&d.bLengthChange)g=ob(a);else if("f"==j&&d.bFilter)g=pb(a);else if("r"==j&&d.bProcessing)g=qb(a);else if("t"==j)g=rb(a);else if("i"==j&&d.bInfo)g=sb(a);else if("p"==
j&&d.bPaginate)g=tb(a);else if(0!==m.ext.feature.length){i=m.ext.feature;q=0;for(n=i.length;q<n;q++)if(j==i[q].cFeature){g=i[q].fnInit(a);break}}g&&(i=a.aanFeatures,i[j]||(i[j]=[]),i[j].push(g),e.append(g))}c.replaceWith(e);a.nHolding=null}function da(a,b){var c=h(b).children("tr"),d,e,f,g,j,i,n,l,q,t;a.splice(0,a.length);f=0;for(i=c.length;f<i;f++)a.push([]);f=0;for(i=c.length;f<i;f++){d=c[f];for(e=d.firstChild;e;){if("TD"==e.nodeName.toUpperCase()||"TH"==e.nodeName.toUpperCase()){l=1*e.getAttribute("colspan");
q=1*e.getAttribute("rowspan");l=!l||0===l||1===l?1:l;q=!q||0===q||1===q?1:q;g=0;for(j=a[f];j[g];)g++;n=g;t=1===l?!0:!1;for(j=0;j<l;j++)for(g=0;g<q;g++)a[f+g][n+j]={cell:e,unique:t},a[f+g].nTr=d}e=e.nextSibling}}}function qa(a,b,c){var d=[];c||(c=a.aoHeader,b&&(c=[],da(c,b)));for(var b=0,e=c.length;b<e;b++)for(var f=0,g=c[b].length;f<g;f++)if(c[b][f].unique&&(!d[f]||!a.bSortCellsTop))d[f]=c[b][f].cell;return d}function ra(a,b,c){u(a,"aoServerParams","serverParams",[b]);if(b&&h.isArray(b)){var d={},
e=/(.*?)\[\]$/;h.each(b,function(a,b){var c=b.name.match(e);c?(c=c[0],d[c]||(d[c]=[]),d[c].push(b.value)):d[b.name]=b.value});b=d}var f,g=a.ajax,j=a.oInstance,i=function(b){u(a,null,"xhr",[a,b,a.jqXHR]);c(b)};if(h.isPlainObject(g)&&g.data){f=g.data;var n=h.isFunction(f)?f(b,a):f,b=h.isFunction(f)&&n?n:h.extend(!0,b,n);delete g.data}n={data:b,success:function(b){var c=b.error||b.sError;c&&L(a,0,c);a.json=b;i(b)},dataType:"json",cache:!1,type:a.sServerMethod,error:function(b,c){var d=u(a,null,"xhr",
[a,null,a.jqXHR]);-1===h.inArray(!0,d)&&("parsererror"==c?L(a,0,"Invalid JSON response",1):4===b.readyState&&L(a,0,"Ajax error",7));C(a,!1)}};a.oAjaxData=b;u(a,null,"preXhr",[a,b]);a.fnServerData?a.fnServerData.call(j,a.sAjaxSource,h.map(b,function(a,b){return{name:b,value:a}}),i,a):a.sAjaxSource||"string"===typeof g?a.jqXHR=h.ajax(h.extend(n,{url:g||a.sAjaxSource})):h.isFunction(g)?a.jqXHR=g.call(j,b,i,a):(a.jqXHR=h.ajax(h.extend(n,g)),g.data=f)}function lb(a){return a.bAjaxDataGet?(a.iDraw++,C(a,
!0),ra(a,ub(a),function(b){vb(a,b)}),!1):!0}function ub(a){var b=a.aoColumns,c=b.length,d=a.oFeatures,e=a.oPreviousSearch,f=a.aoPreSearchCols,g,j=[],i,n,l,q=V(a);g=a._iDisplayStart;i=!1!==d.bPaginate?a._iDisplayLength:-1;var k=function(a,b){j.push({name:a,value:b})};k("sEcho",a.iDraw);k("iColumns",c);k("sColumns",G(b,"sName").join(","));k("iDisplayStart",g);k("iDisplayLength",i);var S={draw:a.iDraw,columns:[],order:[],start:g,length:i,search:{value:e.sSearch,regex:e.bRegex}};for(g=0;g<c;g++)n=b[g],
l=f[g],i="function"==typeof n.mData?"function":n.mData,S.columns.push({data:i,name:n.sName,searchable:n.bSearchable,orderable:n.bSortable,search:{value:l.sSearch,regex:l.bRegex}}),k("mDataProp_"+g,i),d.bFilter&&(k("sSearch_"+g,l.sSearch),k("bRegex_"+g,l.bRegex),k("bSearchable_"+g,n.bSearchable)),d.bSort&&k("bSortable_"+g,n.bSortable);d.bFilter&&(k("sSearch",e.sSearch),k("bRegex",e.bRegex));d.bSort&&(h.each(q,function(a,b){S.order.push({column:b.col,dir:b.dir});k("iSortCol_"+a,b.col);k("sSortDir_"+
a,b.dir)}),k("iSortingCols",q.length));b=m.ext.legacy.ajax;return null===b?a.sAjaxSource?j:S:b?j:S}function vb(a,b){var c=sa(a,b),d=b.sEcho!==k?b.sEcho:b.draw,e=b.iTotalRecords!==k?b.iTotalRecords:b.recordsTotal,f=b.iTotalDisplayRecords!==k?b.iTotalDisplayRecords:b.recordsFiltered;if(d){if(1*d<a.iDraw)return;a.iDraw=1*d}na(a);a._iRecordsTotal=parseInt(e,10);a._iRecordsDisplay=parseInt(f,10);d=0;for(e=c.length;d<e;d++)N(a,c[d]);a.aiDisplay=a.aiDisplayMaster.slice();a.bAjaxDataGet=!1;O(a);a._bInitComplete||
ta(a,b);a.bAjaxDataGet=!0;C(a,!1)}function sa(a,b){var c=h.isPlainObject(a.ajax)&&a.ajax.dataSrc!==k?a.ajax.dataSrc:a.sAjaxDataProp;return"data"===c?b.aaData||b[c]:""!==c?Q(c)(b):b}function pb(a){var b=a.oClasses,c=a.sTableId,d=a.oLanguage,e=a.oPreviousSearch,f=a.aanFeatures,g='<input type="search" class="'+b.sFilterInput+'"/>',j=d.sSearch,j=j.match(/_INPUT_/)?j.replace("_INPUT_",g):j+g,b=h("<div/>",{id:!f.f?c+"_filter":null,"class":b.sFilter}).append(h("<label/>").append(j)),f=function(){var b=!this.value?
"":this.value;b!=e.sSearch&&(fa(a,{sSearch:b,bRegex:e.bRegex,bSmart:e.bSmart,bCaseInsensitive:e.bCaseInsensitive}),a._iDisplayStart=0,O(a))},g=null!==a.searchDelay?a.searchDelay:"ssp"===y(a)?400:0,i=h("input",b).val(e.sSearch).attr("placeholder",d.sSearchPlaceholder).bind("keyup.DT search.DT input.DT paste.DT cut.DT",g?Oa(f,g):f).bind("keypress.DT",function(a){if(13==a.keyCode)return!1}).attr("aria-controls",c);h(a.nTable).on("search.dt.DT",function(b,c){if(a===c)try{i[0]!==I.activeElement&&i.val(e.sSearch)}catch(d){}});
return b[0]}function fa(a,b,c){var d=a.oPreviousSearch,e=a.aoPreSearchCols,f=function(a){d.sSearch=a.sSearch;d.bRegex=a.bRegex;d.bSmart=a.bSmart;d.bCaseInsensitive=a.bCaseInsensitive};Ga(a);if("ssp"!=y(a)){wb(a,b.sSearch,c,b.bEscapeRegex!==k?!b.bEscapeRegex:b.bRegex,b.bSmart,b.bCaseInsensitive);f(b);for(b=0;b<e.length;b++)xb(a,e[b].sSearch,b,e[b].bEscapeRegex!==k?!e[b].bEscapeRegex:e[b].bRegex,e[b].bSmart,e[b].bCaseInsensitive);yb(a)}else f(b);a.bFiltered=!0;u(a,null,"search",[a])}function yb(a){for(var b=
m.ext.search,c=a.aiDisplay,d,e,f=0,g=b.length;f<g;f++){for(var j=[],i=0,n=c.length;i<n;i++)e=c[i],d=a.aoData[e],b[f](a,d._aFilterData,e,d._aData,i)&&j.push(e);c.length=0;h.merge(c,j)}}function xb(a,b,c,d,e,f){if(""!==b)for(var g=a.aiDisplay,d=Pa(b,d,e,f),e=g.length-1;0<=e;e--)b=a.aoData[g[e]]._aFilterData[c],d.test(b)||g.splice(e,1)}function wb(a,b,c,d,e,f){var d=Pa(b,d,e,f),e=a.oPreviousSearch.sSearch,f=a.aiDisplayMaster,g;0!==m.ext.search.length&&(c=!0);g=zb(a);if(0>=b.length)a.aiDisplay=f.slice();
else{if(g||c||e.length>b.length||0!==b.indexOf(e)||a.bSorted)a.aiDisplay=f.slice();b=a.aiDisplay;for(c=b.length-1;0<=c;c--)d.test(a.aoData[b[c]]._sFilterRow)||b.splice(c,1)}}function Pa(a,b,c,d){a=b?a:Qa(a);c&&(a="^(?=.*?"+h.map(a.match(/"[^"]+"|[^ ]+/g)||[""],function(a){if('"'===a.charAt(0))var b=a.match(/^"(.*)"$/),a=b?b[1]:a;return a.replace('"',"")}).join(")(?=.*?")+").*$");return RegExp(a,d?"i":"")}function zb(a){var b=a.aoColumns,c,d,e,f,g,j,i,h,l=m.ext.type.search;c=!1;d=0;for(f=a.aoData.length;d<
f;d++)if(h=a.aoData[d],!h._aFilterData){j=[];e=0;for(g=b.length;e<g;e++)c=b[e],c.bSearchable?(i=B(a,d,e,"filter"),l[c.sType]&&(i=l[c.sType](i)),null===i&&(i=""),"string"!==typeof i&&i.toString&&(i=i.toString())):i="",i.indexOf&&-1!==i.indexOf("&")&&(ua.innerHTML=i,i=Zb?ua.textContent:ua.innerText),i.replace&&(i=i.replace(/[\r\n]/g,"")),j.push(i);h._aFilterData=j;h._sFilterRow=j.join("  ");c=!0}return c}function Ab(a){return{search:a.sSearch,smart:a.bSmart,regex:a.bRegex,caseInsensitive:a.bCaseInsensitive}}
function Bb(a){return{sSearch:a.search,bSmart:a.smart,bRegex:a.regex,bCaseInsensitive:a.caseInsensitive}}function sb(a){var b=a.sTableId,c=a.aanFeatures.i,d=h("<div/>",{"class":a.oClasses.sInfo,id:!c?b+"_info":null});c||(a.aoDrawCallback.push({fn:Cb,sName:"information"}),d.attr("role","status").attr("aria-live","polite"),h(a.nTable).attr("aria-describedby",b+"_info"));return d[0]}function Cb(a){var b=a.aanFeatures.i;if(0!==b.length){var c=a.oLanguage,d=a._iDisplayStart+1,e=a.fnDisplayEnd(),f=a.fnRecordsTotal(),
g=a.fnRecordsDisplay(),j=g?c.sInfo:c.sInfoEmpty;g!==f&&(j+=" "+c.sInfoFiltered);j+=c.sInfoPostFix;j=Db(a,j);c=c.fnInfoCallback;null!==c&&(j=c.call(a.oInstance,a,d,e,f,g,j));h(b).html(j)}}function Db(a,b){var c=a.fnFormatNumber,d=a._iDisplayStart+1,e=a._iDisplayLength,f=a.fnRecordsDisplay(),g=-1===e;return b.replace(/_START_/g,c.call(a,d)).replace(/_END_/g,c.call(a,a.fnDisplayEnd())).replace(/_MAX_/g,c.call(a,a.fnRecordsTotal())).replace(/_TOTAL_/g,c.call(a,f)).replace(/_PAGE_/g,c.call(a,g?1:Math.ceil(d/
e))).replace(/_PAGES_/g,c.call(a,g?1:Math.ceil(f/e)))}function ga(a){var b,c,d=a.iInitDisplayStart,e=a.aoColumns,f;c=a.oFeatures;var g=a.bDeferLoading;if(a.bInitialised){nb(a);kb(a);ea(a,a.aoHeader);ea(a,a.aoFooter);C(a,!0);c.bAutoWidth&&Fa(a);b=0;for(c=e.length;b<c;b++)f=e[b],f.sWidth&&(f.nTh.style.width=x(f.sWidth));u(a,null,"preInit",[a]);T(a);e=y(a);if("ssp"!=e||g)"ajax"==e?ra(a,[],function(c){var f=sa(a,c);for(b=0;b<f.length;b++)N(a,f[b]);a.iInitDisplayStart=d;T(a);C(a,!1);ta(a,c)},a):(C(a,!1),
ta(a))}else setTimeout(function(){ga(a)},200)}function ta(a,b){a._bInitComplete=!0;(b||a.oInit.aaData)&&Y(a);u(a,null,"plugin-init",[a,b]);u(a,"aoInitComplete","init",[a,b])}function Ra(a,b){var c=parseInt(b,10);a._iDisplayLength=c;Sa(a);u(a,null,"length",[a,c])}function ob(a){for(var b=a.oClasses,c=a.sTableId,d=a.aLengthMenu,e=h.isArray(d[0]),f=e?d[0]:d,d=e?d[1]:d,e=h("<select/>",{name:c+"_length","aria-controls":c,"class":b.sLengthSelect}),g=0,j=f.length;g<j;g++)e[0][g]=new Option(d[g],f[g]);var i=
h("<div><label/></div>").addClass(b.sLength);a.aanFeatures.l||(i[0].id=c+"_length");i.children().append(a.oLanguage.sLengthMenu.replace("_MENU_",e[0].outerHTML));h("select",i).val(a._iDisplayLength).bind("change.DT",function(){Ra(a,h(this).val());O(a)});h(a.nTable).bind("length.dt.DT",function(b,c,d){a===c&&h("select",i).val(d)});return i[0]}function tb(a){var b=a.sPaginationType,c=m.ext.pager[b],d="function"===typeof c,e=function(a){O(a)},b=h("<div/>").addClass(a.oClasses.sPaging+b)[0],f=a.aanFeatures;
d||c.fnInit(a,b,e);f.p||(b.id=a.sTableId+"_paginate",a.aoDrawCallback.push({fn:function(a){if(d){var b=a._iDisplayStart,i=a._iDisplayLength,h=a.fnRecordsDisplay(),l=-1===i,b=l?0:Math.ceil(b/i),i=l?1:Math.ceil(h/i),h=c(b,i),k,l=0;for(k=f.p.length;l<k;l++)Na(a,"pageButton")(a,f.p[l],l,h,b,i)}else c.fnUpdate(a,e)},sName:"pagination"}));return b}function Ta(a,b,c){var d=a._iDisplayStart,e=a._iDisplayLength,f=a.fnRecordsDisplay();0===f||-1===e?d=0:"number"===typeof b?(d=b*e,d>f&&(d=0)):"first"==b?d=0:
"previous"==b?(d=0<=e?d-e:0,0>d&&(d=0)):"next"==b?d+e<f&&(d+=e):"last"==b?d=Math.floor((f-1)/e)*e:L(a,0,"Unknown paging action: "+b,5);b=a._iDisplayStart!==d;a._iDisplayStart=d;b&&(u(a,null,"page",[a]),c&&O(a));return b}function qb(a){return h("<div/>",{id:!a.aanFeatures.r?a.sTableId+"_processing":null,"class":a.oClasses.sProcessing}).html(a.oLanguage.sProcessing).insertBefore(a.nTable)[0]}function C(a,b){a.oFeatures.bProcessing&&h(a.aanFeatures.r).css("display",b?"block":"none");u(a,null,"processing",
[a,b])}function rb(a){var b=h(a.nTable);b.attr("role","grid");var c=a.oScroll;if(""===c.sX&&""===c.sY)return a.nTable;var d=c.sX,e=c.sY,f=a.oClasses,g=b.children("caption"),j=g.length?g[0]._captionSide:null,i=h(b[0].cloneNode(!1)),n=h(b[0].cloneNode(!1)),l=b.children("tfoot");l.length||(l=null);i=h("<div/>",{"class":f.sScrollWrapper}).append(h("<div/>",{"class":f.sScrollHead}).css({overflow:"hidden",position:"relative",border:0,width:d?!d?null:x(d):"100%"}).append(h("<div/>",{"class":f.sScrollHeadInner}).css({"box-sizing":"content-box",
width:c.sXInner||"100%"}).append(i.removeAttr("id").css("margin-left",0).append("top"===j?g:null).append(b.children("thead"))))).append(h("<div/>",{"class":f.sScrollBody}).css({position:"relative",overflow:"auto",width:!d?null:x(d)}).append(b));l&&i.append(h("<div/>",{"class":f.sScrollFoot}).css({overflow:"hidden",border:0,width:d?!d?null:x(d):"100%"}).append(h("<div/>",{"class":f.sScrollFootInner}).append(n.removeAttr("id").css("margin-left",0).append("bottom"===j?g:null).append(b.children("tfoot")))));
var b=i.children(),k=b[0],f=b[1],t=l?b[2]:null;if(d)h(f).on("scroll.DT",function(){var a=this.scrollLeft;k.scrollLeft=a;l&&(t.scrollLeft=a)});h(f).css(e&&c.bCollapse?"max-height":"height",e);a.nScrollHead=k;a.nScrollBody=f;a.nScrollFoot=t;a.aoDrawCallback.push({fn:ka,sName:"scrolling"});return i[0]}function ka(a){var b=a.oScroll,c=b.sX,d=b.sXInner,e=b.sY,b=b.iBarWidth,f=h(a.nScrollHead),g=f[0].style,j=f.children("div"),i=j[0].style,n=j.children("table"),j=a.nScrollBody,l=h(j),q=j.style,t=h(a.nScrollFoot).children("div"),
m=t.children("table"),o=h(a.nTHead),F=h(a.nTable),p=F[0],r=p.style,u=a.nTFoot?h(a.nTFoot):null,Eb=a.oBrowser,Ua=Eb.bScrollOversize,s=G(a.aoColumns,"nTh"),P,v,w,y,z=[],A=[],B=[],C=[],D,E=function(a){a=a.style;a.paddingTop="0";a.paddingBottom="0";a.borderTopWidth="0";a.borderBottomWidth="0";a.height=0};v=j.scrollHeight>j.clientHeight;if(a.scrollBarVis!==v&&a.scrollBarVis!==k)a.scrollBarVis=v,Y(a);else{a.scrollBarVis=v;F.children("thead, tfoot").remove();u&&(w=u.clone().prependTo(F),P=u.find("tr"),w=
w.find("tr"));y=o.clone().prependTo(F);o=o.find("tr");v=y.find("tr");y.find("th, td").removeAttr("tabindex");c||(q.width="100%",f[0].style.width="100%");h.each(qa(a,y),function(b,c){D=Z(a,b);c.style.width=a.aoColumns[D].sWidth});u&&J(function(a){a.style.width=""},w);f=F.outerWidth();if(""===c){r.width="100%";if(Ua&&(F.find("tbody").height()>j.offsetHeight||"scroll"==l.css("overflow-y")))r.width=x(F.outerWidth()-b);f=F.outerWidth()}else""!==d&&(r.width=x(d),f=F.outerWidth());J(E,v);J(function(a){B.push(a.innerHTML);
z.push(x(h(a).css("width")))},v);J(function(a,b){if(h.inArray(a,s)!==-1)a.style.width=z[b]},o);h(v).height(0);u&&(J(E,w),J(function(a){C.push(a.innerHTML);A.push(x(h(a).css("width")))},w),J(function(a,b){a.style.width=A[b]},P),h(w).height(0));J(function(a,b){a.innerHTML='<div class="dataTables_sizing" style="height:0;overflow:hidden;">'+B[b]+"</div>";a.style.width=z[b]},v);u&&J(function(a,b){a.innerHTML='<div class="dataTables_sizing" style="height:0;overflow:hidden;">'+C[b]+"</div>";a.style.width=
A[b]},w);if(F.outerWidth()<f){P=j.scrollHeight>j.offsetHeight||"scroll"==l.css("overflow-y")?f+b:f;if(Ua&&(j.scrollHeight>j.offsetHeight||"scroll"==l.css("overflow-y")))r.width=x(P-b);(""===c||""!==d)&&L(a,1,"Possible column misalignment",6)}else P="100%";q.width=x(P);g.width=x(P);u&&(a.nScrollFoot.style.width=x(P));!e&&Ua&&(q.height=x(p.offsetHeight+b));c=F.outerWidth();n[0].style.width=x(c);i.width=x(c);d=F.height()>j.clientHeight||"scroll"==l.css("overflow-y");e="padding"+(Eb.bScrollbarLeft?"Left":
"Right");i[e]=d?b+"px":"0px";u&&(m[0].style.width=x(c),t[0].style.width=x(c),t[0].style[e]=d?b+"px":"0px");F.children("colgroup").insertBefore(F.children("thead"));l.scroll();if((a.bSorted||a.bFiltered)&&!a._drawHold)j.scrollTop=0}}function J(a,b,c){for(var d=0,e=0,f=b.length,g,j;e<f;){g=b[e].firstChild;for(j=c?c[e].firstChild:null;g;)1===g.nodeType&&(c?a(g,j,d):a(g,d),d++),g=g.nextSibling,j=c?j.nextSibling:null;e++}}function Fa(a){var b=a.nTable,c=a.aoColumns,d=a.oScroll,e=d.sY,f=d.sX,g=d.sXInner,
j=c.length,i=la(a,"bVisible"),n=h("th",a.nTHead),l=b.getAttribute("width"),k=b.parentNode,t=!1,m,o,p=a.oBrowser,d=p.bScrollOversize;(m=b.style.width)&&-1!==m.indexOf("%")&&(l=m);for(m=0;m<i.length;m++)o=c[i[m]],null!==o.sWidth&&(o.sWidth=Fb(o.sWidthOrig,k),t=!0);if(d||!t&&!f&&!e&&j==aa(a)&&j==n.length)for(m=0;m<j;m++)i=Z(a,m),null!==i&&(c[i].sWidth=x(n.eq(m).width()));else{j=h(b).clone().css("visibility","hidden").removeAttr("id");j.find("tbody tr").remove();var r=h("<tr/>").appendTo(j.find("tbody"));
j.find("thead, tfoot").remove();j.append(h(a.nTHead).clone()).append(h(a.nTFoot).clone());j.find("tfoot th, tfoot td").css("width","");n=qa(a,j.find("thead")[0]);for(m=0;m<i.length;m++)o=c[i[m]],n[m].style.width=null!==o.sWidthOrig&&""!==o.sWidthOrig?x(o.sWidthOrig):"",o.sWidthOrig&&f&&h(n[m]).append(h("<div/>").css({width:o.sWidthOrig,margin:0,padding:0,border:0,height:1}));if(a.aoData.length)for(m=0;m<i.length;m++)t=i[m],o=c[t],h(Gb(a,t)).clone(!1).append(o.sContentPadding).appendTo(r);h("[name]",
j).removeAttr("name");o=h("<div/>").css(f||e?{position:"absolute",top:0,left:0,height:1,right:0,overflow:"hidden"}:{}).append(j).appendTo(k);f&&g?j.width(g):f?(j.css("width","auto"),j.removeAttr("width"),j.width()<k.clientWidth&&l&&j.width(k.clientWidth)):e?j.width(k.clientWidth):l&&j.width(l);for(m=e=0;m<i.length;m++)k=h(n[m]),g=k.outerWidth()-k.width(),k=p.bBounding?Math.ceil(n[m].getBoundingClientRect().width):k.outerWidth(),e+=k,c[i[m]].sWidth=x(k-g);b.style.width=x(e);o.remove()}l&&(b.style.width=
x(l));if((l||f)&&!a._reszEvt)b=function(){h(D).bind("resize.DT-"+a.sInstance,Oa(function(){Y(a)}))},d?setTimeout(b,1E3):b(),a._reszEvt=!0}function Fb(a,b){if(!a)return 0;var c=h("<div/>").css("width",x(a)).appendTo(b||I.body),d=c[0].offsetWidth;c.remove();return d}function Gb(a,b){var c=Hb(a,b);if(0>c)return null;var d=a.aoData[c];return!d.nTr?h("<td/>").html(B(a,c,b,"display"))[0]:d.anCells[b]}function Hb(a,b){for(var c,d=-1,e=-1,f=0,g=a.aoData.length;f<g;f++)c=B(a,f,b,"display")+"",c=c.replace($b,
""),c=c.replace(/&nbsp;/g," "),c.length>d&&(d=c.length,e=f);return e}function x(a){return null===a?"0px":"number"==typeof a?0>a?"0px":a+"px":a.match(/\d$/)?a+"px":a}function V(a){var b,c,d=[],e=a.aoColumns,f,g,j,i;b=a.aaSortingFixed;c=h.isPlainObject(b);var n=[];f=function(a){a.length&&!h.isArray(a[0])?n.push(a):h.merge(n,a)};h.isArray(b)&&f(b);c&&b.pre&&f(b.pre);f(a.aaSorting);c&&b.post&&f(b.post);for(a=0;a<n.length;a++){i=n[a][0];f=e[i].aDataSort;b=0;for(c=f.length;b<c;b++)g=f[b],j=e[g].sType||
"string",n[a]._idx===k&&(n[a]._idx=h.inArray(n[a][1],e[g].asSorting)),d.push({src:i,col:g,dir:n[a][1],index:n[a]._idx,type:j,formatter:m.ext.type.order[j+"-pre"]})}return d}function mb(a){var b,c,d=[],e=m.ext.type.order,f=a.aoData,g=0,j,i=a.aiDisplayMaster,h;Ga(a);h=V(a);b=0;for(c=h.length;b<c;b++)j=h[b],j.formatter&&g++,Ib(a,j.col);if("ssp"!=y(a)&&0!==h.length){b=0;for(c=i.length;b<c;b++)d[i[b]]=b;g===h.length?i.sort(function(a,b){var c,e,g,j,i=h.length,k=f[a]._aSortData,m=f[b]._aSortData;for(g=
0;g<i;g++)if(j=h[g],c=k[j.col],e=m[j.col],c=c<e?-1:c>e?1:0,0!==c)return"asc"===j.dir?c:-c;c=d[a];e=d[b];return c<e?-1:c>e?1:0}):i.sort(function(a,b){var c,g,j,i,k=h.length,m=f[a]._aSortData,p=f[b]._aSortData;for(j=0;j<k;j++)if(i=h[j],c=m[i.col],g=p[i.col],i=e[i.type+"-"+i.dir]||e["string-"+i.dir],c=i(c,g),0!==c)return c;c=d[a];g=d[b];return c<g?-1:c>g?1:0})}a.bSorted=!0}function Jb(a){for(var b,c,d=a.aoColumns,e=V(a),a=a.oLanguage.oAria,f=0,g=d.length;f<g;f++){c=d[f];var j=c.asSorting;b=c.sTitle.replace(/<.*?>/g,
"");var i=c.nTh;i.removeAttribute("aria-sort");c.bSortable&&(0<e.length&&e[0].col==f?(i.setAttribute("aria-sort","asc"==e[0].dir?"ascending":"descending"),c=j[e[0].index+1]||j[0]):c=j[0],b+="asc"===c?a.sSortAscending:a.sSortDescending);i.setAttribute("aria-label",b)}}function Va(a,b,c,d){var e=a.aaSorting,f=a.aoColumns[b].asSorting,g=function(a,b){var c=a._idx;c===k&&(c=h.inArray(a[1],f));return c+1<f.length?c+1:b?null:0};"number"===typeof e[0]&&(e=a.aaSorting=[e]);c&&a.oFeatures.bSortMulti?(c=h.inArray(b,
G(e,"0")),-1!==c?(b=g(e[c],!0),null===b&&1===e.length&&(b=0),null===b?e.splice(c,1):(e[c][1]=f[b],e[c]._idx=b)):(e.push([b,f[0],0]),e[e.length-1]._idx=0)):e.length&&e[0][0]==b?(b=g(e[0]),e.length=1,e[0][1]=f[b],e[0]._idx=b):(e.length=0,e.push([b,f[0]]),e[0]._idx=0);T(a);"function"==typeof d&&d(a)}function Ma(a,b,c,d){var e=a.aoColumns[c];Wa(b,{},function(b){!1!==e.bSortable&&(a.oFeatures.bProcessing?(C(a,!0),setTimeout(function(){Va(a,c,b.shiftKey,d);"ssp"!==y(a)&&C(a,!1)},0)):Va(a,c,b.shiftKey,d))})}
function va(a){var b=a.aLastSort,c=a.oClasses.sSortColumn,d=V(a),e=a.oFeatures,f,g;if(e.bSort&&e.bSortClasses){e=0;for(f=b.length;e<f;e++)g=b[e].src,h(G(a.aoData,"anCells",g)).removeClass(c+(2>e?e+1:3));e=0;for(f=d.length;e<f;e++)g=d[e].src,h(G(a.aoData,"anCells",g)).addClass(c+(2>e?e+1:3))}a.aLastSort=d}function Ib(a,b){var c=a.aoColumns[b],d=m.ext.order[c.sSortDataType],e;d&&(e=d.call(a.oInstance,a,b,$(a,b)));for(var f,g=m.ext.type.order[c.sType+"-pre"],j=0,i=a.aoData.length;j<i;j++)if(c=a.aoData[j],
c._aSortData||(c._aSortData=[]),!c._aSortData[b]||d)f=d?e[j]:B(a,j,b,"sort"),c._aSortData[b]=g?g(f):f}function wa(a){if(a.oFeatures.bStateSave&&!a.bDestroying){var b={time:+new Date,start:a._iDisplayStart,length:a._iDisplayLength,order:h.extend(!0,[],a.aaSorting),search:Ab(a.oPreviousSearch),columns:h.map(a.aoColumns,function(b,d){return{visible:b.bVisible,search:Ab(a.aoPreSearchCols[d])}})};u(a,"aoStateSaveParams","stateSaveParams",[a,b]);a.oSavedState=b;a.fnStateSaveCallback.call(a.oInstance,a,
b)}}function Kb(a){var b,c,d=a.aoColumns;if(a.oFeatures.bStateSave){var e=a.fnStateLoadCallback.call(a.oInstance,a);if(e&&e.time&&(b=u(a,"aoStateLoadParams","stateLoadParams",[a,e]),-1===h.inArray(!1,b)&&(b=a.iStateDuration,!(0<b&&e.time<+new Date-1E3*b)&&d.length===e.columns.length))){a.oLoadedState=h.extend(!0,{},e);e.start!==k&&(a._iDisplayStart=e.start,a.iInitDisplayStart=e.start);e.length!==k&&(a._iDisplayLength=e.length);e.order!==k&&(a.aaSorting=[],h.each(e.order,function(b,c){a.aaSorting.push(c[0]>=
d.length?[0,c[1]]:c)}));e.search!==k&&h.extend(a.oPreviousSearch,Bb(e.search));b=0;for(c=e.columns.length;b<c;b++){var f=e.columns[b];f.visible!==k&&(d[b].bVisible=f.visible);f.search!==k&&h.extend(a.aoPreSearchCols[b],Bb(f.search))}u(a,"aoStateLoaded","stateLoaded",[a,e])}}}function xa(a){var b=m.settings,a=h.inArray(a,G(b,"nTable"));return-1!==a?b[a]:null}function L(a,b,c,d){c="DataTables warning: "+(a?"table id="+a.sTableId+" - ":"")+c;d&&(c+=". For more information about this error, please see http://datatables.net/tn/"+
d);if(b)D.console&&console.log&&console.log(c);else if(b=m.ext,b=b.sErrMode||b.errMode,a&&u(a,null,"error",[a,d,c]),"alert"==b)alert(c);else{if("throw"==b)throw Error(c);"function"==typeof b&&b(a,d,c)}}function E(a,b,c,d){h.isArray(c)?h.each(c,function(c,d){h.isArray(d)?E(a,b,d[0],d[1]):E(a,b,d)}):(d===k&&(d=c),b[c]!==k&&(a[d]=b[c]))}function Lb(a,b,c){var d,e;for(e in b)b.hasOwnProperty(e)&&(d=b[e],h.isPlainObject(d)?(h.isPlainObject(a[e])||(a[e]={}),h.extend(!0,a[e],d)):a[e]=c&&"data"!==e&&"aaData"!==
e&&h.isArray(d)?d.slice():d);return a}function Wa(a,b,c){h(a).bind("click.DT",b,function(b){a.blur();c(b)}).bind("keypress.DT",b,function(a){13===a.which&&(a.preventDefault(),c(a))}).bind("selectstart.DT",function(){return!1})}function z(a,b,c,d){c&&a[b].push({fn:c,sName:d})}function u(a,b,c,d){var e=[];b&&(e=h.map(a[b].slice().reverse(),function(b){return b.fn.apply(a.oInstance,d)}));null!==c&&(b=h.Event(c+".dt"),h(a.nTable).trigger(b,d),e.push(b.result));return e}function Sa(a){var b=a._iDisplayStart,
c=a.fnDisplayEnd(),d=a._iDisplayLength;b>=c&&(b=c-d);b-=b%d;if(-1===d||0>b)b=0;a._iDisplayStart=b}function Na(a,b){var c=a.renderer,d=m.ext.renderer[b];return h.isPlainObject(c)&&c[b]?d[c[b]]||d._:"string"===typeof c?d[c]||d._:d._}function y(a){return a.oFeatures.bServerSide?"ssp":a.ajax||a.sAjaxSource?"ajax":"dom"}function ya(a,b){var c=[],c=Mb.numbers_length,d=Math.floor(c/2);b<=c?c=W(0,b):a<=d?(c=W(0,c-2),c.push("ellipsis"),c.push(b-1)):(a>=b-1-d?c=W(b-(c-2),b):(c=W(a-d+2,a+d-1),c.push("ellipsis"),
c.push(b-1)),c.splice(0,0,"ellipsis"),c.splice(0,0,0));c.DT_el="span";return c}function db(a){h.each({num:function(b){return za(b,a)},"num-fmt":function(b){return za(b,a,Xa)},"html-num":function(b){return za(b,a,Aa)},"html-num-fmt":function(b){return za(b,a,Aa,Xa)}},function(b,c){v.type.order[b+a+"-pre"]=c;b.match(/^html\-/)&&(v.type.search[b+a]=v.type.search.html)})}function Nb(a){return function(){var b=[xa(this[m.ext.iApiIndex])].concat(Array.prototype.slice.call(arguments));return m.ext.internal[a].apply(this,
b)}}var m=function(a){this.$=function(a,b){return this.api(!0).$(a,b)};this._=function(a,b){return this.api(!0).rows(a,b).data()};this.api=function(a){return a?new r(xa(this[v.iApiIndex])):new r(this)};this.fnAddData=function(a,b){var c=this.api(!0),d=h.isArray(a)&&(h.isArray(a[0])||h.isPlainObject(a[0]))?c.rows.add(a):c.row.add(a);(b===k||b)&&c.draw();return d.flatten().toArray()};this.fnAdjustColumnSizing=function(a){var b=this.api(!0).columns.adjust(),c=b.settings()[0],d=c.oScroll;a===k||a?b.draw(!1):
(""!==d.sX||""!==d.sY)&&ka(c)};this.fnClearTable=function(a){var b=this.api(!0).clear();(a===k||a)&&b.draw()};this.fnClose=function(a){this.api(!0).row(a).child.hide()};this.fnDeleteRow=function(a,b,c){var d=this.api(!0),a=d.rows(a),e=a.settings()[0],h=e.aoData[a[0][0]];a.remove();b&&b.call(this,e,h);(c===k||c)&&d.draw();return h};this.fnDestroy=function(a){this.api(!0).destroy(a)};this.fnDraw=function(a){this.api(!0).draw(a)};this.fnFilter=function(a,b,c,d,e,h){e=this.api(!0);null===b||b===k?e.search(a,
c,d,h):e.column(b).search(a,c,d,h);e.draw()};this.fnGetData=function(a,b){var c=this.api(!0);if(a!==k){var d=a.nodeName?a.nodeName.toLowerCase():"";return b!==k||"td"==d||"th"==d?c.cell(a,b).data():c.row(a).data()||null}return c.data().toArray()};this.fnGetNodes=function(a){var b=this.api(!0);return a!==k?b.row(a).node():b.rows().nodes().flatten().toArray()};this.fnGetPosition=function(a){var b=this.api(!0),c=a.nodeName.toUpperCase();return"TR"==c?b.row(a).index():"TD"==c||"TH"==c?(a=b.cell(a).index(),
[a.row,a.columnVisible,a.column]):null};this.fnIsOpen=function(a){return this.api(!0).row(a).child.isShown()};this.fnOpen=function(a,b,c){return this.api(!0).row(a).child(b,c).show().child()[0]};this.fnPageChange=function(a,b){var c=this.api(!0).page(a);(b===k||b)&&c.draw(!1)};this.fnSetColumnVis=function(a,b,c){a=this.api(!0).column(a).visible(b);(c===k||c)&&a.columns.adjust().draw()};this.fnSettings=function(){return xa(this[v.iApiIndex])};this.fnSort=function(a){this.api(!0).order(a).draw()};this.fnSortListener=
function(a,b,c){this.api(!0).order.listener(a,b,c)};this.fnUpdate=function(a,b,c,d,e){var h=this.api(!0);c===k||null===c?h.row(b).data(a):h.cell(b,c).data(a);(e===k||e)&&h.columns.adjust();(d===k||d)&&h.draw();return 0};this.fnVersionCheck=v.fnVersionCheck;var b=this,c=a===k,d=this.length;c&&(a={});this.oApi=this.internal=v.internal;for(var e in m.ext.internal)e&&(this[e]=Nb(e));this.each(function(){var e={},e=1<d?Lb(e,a,!0):a,g=0,j,i=this.getAttribute("id"),n=!1,l=m.defaults,q=h(this);if("table"!=
this.nodeName.toLowerCase())L(null,0,"Non-table node initialisation ("+this.nodeName+")",2);else{eb(l);fb(l.column);K(l,l,!0);K(l.column,l.column,!0);K(l,h.extend(e,q.data()));var t=m.settings,g=0;for(j=t.length;g<j;g++){var p=t[g];if(p.nTable==this||p.nTHead.parentNode==this||p.nTFoot&&p.nTFoot.parentNode==this){g=e.bRetrieve!==k?e.bRetrieve:l.bRetrieve;if(c||g)return p.oInstance;if(e.bDestroy!==k?e.bDestroy:l.bDestroy){p.oInstance.fnDestroy();break}else{L(p,0,"Cannot reinitialise DataTable",3);
return}}if(p.sTableId==this.id){t.splice(g,1);break}}if(null===i||""===i)this.id=i="DataTables_Table_"+m.ext._unique++;var o=h.extend(!0,{},m.models.oSettings,{sDestroyWidth:q[0].style.width,sInstance:i,sTableId:i});o.nTable=this;o.oApi=b.internal;o.oInit=e;t.push(o);o.oInstance=1===b.length?b:q.dataTable();eb(e);e.oLanguage&&Da(e.oLanguage);e.aLengthMenu&&!e.iDisplayLength&&(e.iDisplayLength=h.isArray(e.aLengthMenu[0])?e.aLengthMenu[0][0]:e.aLengthMenu[0]);e=Lb(h.extend(!0,{},l),e);E(o.oFeatures,
e,"bPaginate bLengthChange bFilter bSort bSortMulti bInfo bProcessing bAutoWidth bSortClasses bServerSide bDeferRender".split(" "));E(o,e,["asStripeClasses","ajax","fnServerData","fnFormatNumber","sServerMethod","aaSorting","aaSortingFixed","aLengthMenu","sPaginationType","sAjaxSource","sAjaxDataProp","iStateDuration","sDom","bSortCellsTop","iTabIndex","fnStateLoadCallback","fnStateSaveCallback","renderer","searchDelay","rowId",["iCookieDuration","iStateDuration"],["oSearch","oPreviousSearch"],["aoSearchCols",
"aoPreSearchCols"],["iDisplayLength","_iDisplayLength"],["bJQueryUI","bJUI"]]);E(o.oScroll,e,[["sScrollX","sX"],["sScrollXInner","sXInner"],["sScrollY","sY"],["bScrollCollapse","bCollapse"]]);E(o.oLanguage,e,"fnInfoCallback");z(o,"aoDrawCallback",e.fnDrawCallback,"user");z(o,"aoServerParams",e.fnServerParams,"user");z(o,"aoStateSaveParams",e.fnStateSaveParams,"user");z(o,"aoStateLoadParams",e.fnStateLoadParams,"user");z(o,"aoStateLoaded",e.fnStateLoaded,"user");z(o,"aoRowCallback",e.fnRowCallback,
"user");z(o,"aoRowCreatedCallback",e.fnCreatedRow,"user");z(o,"aoHeaderCallback",e.fnHeaderCallback,"user");z(o,"aoFooterCallback",e.fnFooterCallback,"user");z(o,"aoInitComplete",e.fnInitComplete,"user");z(o,"aoPreDrawCallback",e.fnPreDrawCallback,"user");o.rowIdFn=Q(e.rowId);gb(o);i=o.oClasses;e.bJQueryUI?(h.extend(i,m.ext.oJUIClasses,e.oClasses),e.sDom===l.sDom&&"lfrtip"===l.sDom&&(o.sDom='<"H"lfr>t<"F"ip>'),o.renderer)?h.isPlainObject(o.renderer)&&!o.renderer.header&&(o.renderer.header="jqueryui"):
o.renderer="jqueryui":h.extend(i,m.ext.classes,e.oClasses);q.addClass(i.sTable);o.iInitDisplayStart===k&&(o.iInitDisplayStart=e.iDisplayStart,o._iDisplayStart=e.iDisplayStart);null!==e.iDeferLoading&&(o.bDeferLoading=!0,g=h.isArray(e.iDeferLoading),o._iRecordsDisplay=g?e.iDeferLoading[0]:e.iDeferLoading,o._iRecordsTotal=g?e.iDeferLoading[1]:e.iDeferLoading);var r=o.oLanguage;h.extend(!0,r,e.oLanguage);""!==r.sUrl&&(h.ajax({dataType:"json",url:r.sUrl,success:function(a){Da(a);K(l.oLanguage,a);h.extend(true,
r,a);ga(o)},error:function(){ga(o)}}),n=!0);null===e.asStripeClasses&&(o.asStripeClasses=[i.sStripeOdd,i.sStripeEven]);var g=o.asStripeClasses,v=q.children("tbody").find("tr").eq(0);-1!==h.inArray(!0,h.map(g,function(a){return v.hasClass(a)}))&&(h("tbody tr",this).removeClass(g.join(" ")),o.asDestroyStripes=g.slice());t=[];g=this.getElementsByTagName("thead");0!==g.length&&(da(o.aoHeader,g[0]),t=qa(o));if(null===e.aoColumns){p=[];g=0;for(j=t.length;g<j;g++)p.push(null)}else p=e.aoColumns;g=0;for(j=
p.length;g<j;g++)Ea(o,t?t[g]:null);ib(o,e.aoColumnDefs,p,function(a,b){ja(o,a,b)});if(v.length){var s=function(a,b){return a.getAttribute("data-"+b)!==null?b:null};h(v[0]).children("th, td").each(function(a,b){var c=o.aoColumns[a];if(c.mData===a){var d=s(b,"sort")||s(b,"order"),e=s(b,"filter")||s(b,"search");if(d!==null||e!==null){c.mData={_:a+".display",sort:d!==null?a+".@data-"+d:k,type:d!==null?a+".@data-"+d:k,filter:e!==null?a+".@data-"+e:k};ja(o,a)}}})}var w=o.oFeatures;e.bStateSave&&(w.bStateSave=
!0,Kb(o,e),z(o,"aoDrawCallback",wa,"state_save"));if(e.aaSorting===k){t=o.aaSorting;g=0;for(j=t.length;g<j;g++)t[g][1]=o.aoColumns[g].asSorting[0]}va(o);w.bSort&&z(o,"aoDrawCallback",function(){if(o.bSorted){var a=V(o),b={};h.each(a,function(a,c){b[c.src]=c.dir});u(o,null,"order",[o,a,b]);Jb(o)}});z(o,"aoDrawCallback",function(){(o.bSorted||y(o)==="ssp"||w.bDeferRender)&&va(o)},"sc");g=q.children("caption").each(function(){this._captionSide=q.css("caption-side")});j=q.children("thead");0===j.length&&
(j=h("<thead/>").appendTo(this));o.nTHead=j[0];j=q.children("tbody");0===j.length&&(j=h("<tbody/>").appendTo(this));o.nTBody=j[0];j=q.children("tfoot");if(0===j.length&&0<g.length&&(""!==o.oScroll.sX||""!==o.oScroll.sY))j=h("<tfoot/>").appendTo(this);0===j.length||0===j.children().length?q.addClass(i.sNoFooter):0<j.length&&(o.nTFoot=j[0],da(o.aoFooter,o.nTFoot));if(e.aaData)for(g=0;g<e.aaData.length;g++)N(o,e.aaData[g]);else(o.bDeferLoading||"dom"==y(o))&&ma(o,h(o.nTBody).children("tr"));o.aiDisplay=
o.aiDisplayMaster.slice();o.bInitialised=!0;!1===n&&ga(o)}});b=null;return this},v,r,p,s,Ya={},Ob=/[\r\n]/g,Aa=/<.*?>/g,ac=/^[\w\+\-]/,bc=/[\w\+\-]$/,cc=RegExp("(\\/|\\.|\\*|\\+|\\?|\\||\\(|\\)|\\[|\\]|\\{|\\}|\\\\|\\$|\\^|\\-)","g"),Xa=/[',$£€¥%\u2009\u202F\u20BD\u20a9\u20BArfk]/gi,M=function(a){return!a||!0===a||"-"===a?!0:!1},Pb=function(a){var b=parseInt(a,10);return!isNaN(b)&&isFinite(a)?b:null},Qb=function(a,b){Ya[b]||(Ya[b]=RegExp(Qa(b),"g"));return"string"===typeof a&&"."!==b?a.replace(/\./g,
"").replace(Ya[b],"."):a},Za=function(a,b,c){var d="string"===typeof a;if(M(a))return!0;b&&d&&(a=Qb(a,b));c&&d&&(a=a.replace(Xa,""));return!isNaN(parseFloat(a))&&isFinite(a)},Rb=function(a,b,c){return M(a)?!0:!(M(a)||"string"===typeof a)?null:Za(a.replace(Aa,""),b,c)?!0:null},G=function(a,b,c){var d=[],e=0,f=a.length;if(c!==k)for(;e<f;e++)a[e]&&a[e][b]&&d.push(a[e][b][c]);else for(;e<f;e++)a[e]&&d.push(a[e][b]);return d},ha=function(a,b,c,d){var e=[],f=0,g=b.length;if(d!==k)for(;f<g;f++)a[b[f]][c]&&
e.push(a[b[f]][c][d]);else for(;f<g;f++)e.push(a[b[f]][c]);return e},W=function(a,b){var c=[],d;b===k?(b=0,d=a):(d=b,b=a);for(var e=b;e<d;e++)c.push(e);return c},Sb=function(a){for(var b=[],c=0,d=a.length;c<d;c++)a[c]&&b.push(a[c]);return b},pa=function(a){var b=[],c,d,e=a.length,f,g=0;d=0;a:for(;d<e;d++){c=a[d];for(f=0;f<g;f++)if(b[f]===c)continue a;b.push(c);g++}return b};m.util={throttle:function(a,b){var c=b!==k?b:200,d,e;return function(){var b=this,g=+new Date,h=arguments;d&&g<d+c?(clearTimeout(e),
e=setTimeout(function(){d=k;a.apply(b,h)},c)):(d=g,a.apply(b,h))}},escapeRegex:function(a){return a.replace(cc,"\\$1")}};var A=function(a,b,c){a[b]!==k&&(a[c]=a[b])},ba=/\[.*?\]$/,U=/\(\)$/,Qa=m.util.escapeRegex,ua=h("<div>")[0],Zb=ua.textContent!==k,$b=/<.*?>/g,Oa=m.util.throttle,Tb=[],w=Array.prototype,dc=function(a){var b,c,d=m.settings,e=h.map(d,function(a){return a.nTable});if(a){if(a.nTable&&a.oApi)return[a];if(a.nodeName&&"table"===a.nodeName.toLowerCase())return b=h.inArray(a,e),-1!==b?[d[b]]:
null;if(a&&"function"===typeof a.settings)return a.settings().toArray();"string"===typeof a?c=h(a):a instanceof h&&(c=a)}else return[];if(c)return c.map(function(){b=h.inArray(this,e);return-1!==b?d[b]:null}).toArray()};r=function(a,b){if(!(this instanceof r))return new r(a,b);var c=[],d=function(a){(a=dc(a))&&(c=c.concat(a))};if(h.isArray(a))for(var e=0,f=a.length;e<f;e++)d(a[e]);else d(a);this.context=pa(c);b&&h.merge(this,b);this.selector={rows:null,cols:null,opts:null};r.extend(this,this,Tb)};
m.Api=r;h.extend(r.prototype,{any:function(){return 0!==this.count()},concat:w.concat,context:[],count:function(){return this.flatten().length},each:function(a){for(var b=0,c=this.length;b<c;b++)a.call(this,this[b],b,this);return this},eq:function(a){var b=this.context;return b.length>a?new r(b[a],this[a]):null},filter:function(a){var b=[];if(w.filter)b=w.filter.call(this,a,this);else for(var c=0,d=this.length;c<d;c++)a.call(this,this[c],c,this)&&b.push(this[c]);return new r(this.context,b)},flatten:function(){var a=
[];return new r(this.context,a.concat.apply(a,this.toArray()))},join:w.join,indexOf:w.indexOf||function(a,b){for(var c=b||0,d=this.length;c<d;c++)if(this[c]===a)return c;return-1},iterator:function(a,b,c,d){var e=[],f,g,h,i,n,l=this.context,m,t,p=this.selector;"string"===typeof a&&(d=c,c=b,b=a,a=!1);g=0;for(h=l.length;g<h;g++){var o=new r(l[g]);if("table"===b)f=c.call(o,l[g],g),f!==k&&e.push(f);else if("columns"===b||"rows"===b)f=c.call(o,l[g],this[g],g),f!==k&&e.push(f);else if("column"===b||"column-rows"===
b||"row"===b||"cell"===b){t=this[g];"column-rows"===b&&(m=Ba(l[g],p.opts));i=0;for(n=t.length;i<n;i++)f=t[i],f="cell"===b?c.call(o,l[g],f.row,f.column,g,i):c.call(o,l[g],f,g,i,m),f!==k&&e.push(f)}}return e.length||d?(a=new r(l,a?e.concat.apply([],e):e),b=a.selector,b.rows=p.rows,b.cols=p.cols,b.opts=p.opts,a):this},lastIndexOf:w.lastIndexOf||function(a,b){return this.indexOf.apply(this.toArray.reverse(),arguments)},length:0,map:function(a){var b=[];if(w.map)b=w.map.call(this,a,this);else for(var c=
0,d=this.length;c<d;c++)b.push(a.call(this,this[c],c));return new r(this.context,b)},pluck:function(a){return this.map(function(b){return b[a]})},pop:w.pop,push:w.push,reduce:w.reduce||function(a,b){return hb(this,a,b,0,this.length,1)},reduceRight:w.reduceRight||function(a,b){return hb(this,a,b,this.length-1,-1,-1)},reverse:w.reverse,selector:null,shift:w.shift,sort:w.sort,splice:w.splice,toArray:function(){return w.slice.call(this)},to$:function(){return h(this)},toJQuery:function(){return h(this)},
unique:function(){return new r(this.context,pa(this))},unshift:w.unshift});r.extend=function(a,b,c){if(c.length&&b&&(b instanceof r||b.__dt_wrapper)){var d,e,f,g=function(a,b,c){return function(){var d=b.apply(a,arguments);r.extend(d,d,c.methodExt);return d}};d=0;for(e=c.length;d<e;d++)f=c[d],b[f.name]="function"===typeof f.val?g(a,f.val,f):h.isPlainObject(f.val)?{}:f.val,b[f.name].__dt_wrapper=!0,r.extend(a,b[f.name],f.propExt)}};r.register=p=function(a,b){if(h.isArray(a))for(var c=0,d=a.length;c<
d;c++)r.register(a[c],b);else for(var e=a.split("."),f=Tb,g,j,c=0,d=e.length;c<d;c++){g=(j=-1!==e[c].indexOf("()"))?e[c].replace("()",""):e[c];var i;a:{i=0;for(var n=f.length;i<n;i++)if(f[i].name===g){i=f[i];break a}i=null}i||(i={name:g,val:{},methodExt:[],propExt:[]},f.push(i));c===d-1?i.val=b:f=j?i.methodExt:i.propExt}};r.registerPlural=s=function(a,b,c){r.register(a,c);r.register(b,function(){var a=c.apply(this,arguments);return a===this?this:a instanceof r?a.length?h.isArray(a[0])?new r(a.context,
a[0]):a[0]:k:a})};p("tables()",function(a){var b;if(a){b=r;var c=this.context;if("number"===typeof a)a=[c[a]];else var d=h.map(c,function(a){return a.nTable}),a=h(d).filter(a).map(function(){var a=h.inArray(this,d);return c[a]}).toArray();b=new b(a)}else b=this;return b});p("table()",function(a){var a=this.tables(a),b=a.context;return b.length?new r(b[0]):a});s("tables().nodes()","table().node()",function(){return this.iterator("table",function(a){return a.nTable},1)});s("tables().body()","table().body()",
function(){return this.iterator("table",function(a){return a.nTBody},1)});s("tables().header()","table().header()",function(){return this.iterator("table",function(a){return a.nTHead},1)});s("tables().footer()","table().footer()",function(){return this.iterator("table",function(a){return a.nTFoot},1)});s("tables().containers()","table().container()",function(){return this.iterator("table",function(a){return a.nTableWrapper},1)});p("draw()",function(a){return this.iterator("table",function(b){"page"===
a?O(b):("string"===typeof a&&(a="full-hold"===a?!1:!0),T(b,!1===a))})});p("page()",function(a){return a===k?this.page.info().page:this.iterator("table",function(b){Ta(b,a)})});p("page.info()",function(){if(0===this.context.length)return k;var a=this.context[0],b=a._iDisplayStart,c=a.oFeatures.bPaginate?a._iDisplayLength:-1,d=a.fnRecordsDisplay(),e=-1===c;return{page:e?0:Math.floor(b/c),pages:e?1:Math.ceil(d/c),start:b,end:a.fnDisplayEnd(),length:c,recordsTotal:a.fnRecordsTotal(),recordsDisplay:d,
serverSide:"ssp"===y(a)}});p("page.len()",function(a){return a===k?0!==this.context.length?this.context[0]._iDisplayLength:k:this.iterator("table",function(b){Ra(b,a)})});var Ub=function(a,b,c){if(c){var d=new r(a);d.one("draw",function(){c(d.ajax.json())})}if("ssp"==y(a))T(a,b);else{C(a,!0);var e=a.jqXHR;e&&4!==e.readyState&&e.abort();ra(a,[],function(c){na(a);for(var c=sa(a,c),d=0,e=c.length;d<e;d++)N(a,c[d]);T(a,b);C(a,!1)})}};p("ajax.json()",function(){var a=this.context;if(0<a.length)return a[0].json});
p("ajax.params()",function(){var a=this.context;if(0<a.length)return a[0].oAjaxData});p("ajax.reload()",function(a,b){return this.iterator("table",function(c){Ub(c,!1===b,a)})});p("ajax.url()",function(a){var b=this.context;if(a===k){if(0===b.length)return k;b=b[0];return b.ajax?h.isPlainObject(b.ajax)?b.ajax.url:b.ajax:b.sAjaxSource}return this.iterator("table",function(b){h.isPlainObject(b.ajax)?b.ajax.url=a:b.ajax=a})});p("ajax.url().load()",function(a,b){return this.iterator("table",function(c){Ub(c,
!1===b,a)})});var $a=function(a,b,c,d,e){var f=[],g,j,i,n,l,m;i=typeof b;if(!b||"string"===i||"function"===i||b.length===k)b=[b];i=0;for(n=b.length;i<n;i++){j=b[i]&&b[i].split?b[i].split(","):[b[i]];l=0;for(m=j.length;l<m;l++)(g=c("string"===typeof j[l]?h.trim(j[l]):j[l]))&&g.length&&(f=f.concat(g))}a=v.selector[a];if(a.length){i=0;for(n=a.length;i<n;i++)f=a[i](d,e,f)}return pa(f)},ab=function(a){a||(a={});a.filter&&a.search===k&&(a.search=a.filter);return h.extend({search:"none",order:"current",
page:"all"},a)},bb=function(a){for(var b=0,c=a.length;b<c;b++)if(0<a[b].length)return a[0]=a[b],a[0].length=1,a.length=1,a.context=[a.context[b]],a;a.length=0;return a},Ba=function(a,b){var c,d,e,f=[],g=a.aiDisplay;c=a.aiDisplayMaster;var j=b.search;d=b.order;e=b.page;if("ssp"==y(a))return"removed"===j?[]:W(0,c.length);if("current"==e){c=a._iDisplayStart;for(d=a.fnDisplayEnd();c<d;c++)f.push(g[c])}else if("current"==d||"applied"==d)f="none"==j?c.slice():"applied"==j?g.slice():h.map(c,function(a){return-1===
h.inArray(a,g)?a:null});else if("index"==d||"original"==d){c=0;for(d=a.aoData.length;c<d;c++)"none"==j?f.push(c):(e=h.inArray(c,g),(-1===e&&"removed"==j||0<=e&&"applied"==j)&&f.push(c))}return f};p("rows()",function(a,b){a===k?a="":h.isPlainObject(a)&&(b=a,a="");var b=ab(b),c=this.iterator("table",function(c){var e=b;return $a("row",a,function(a){var b=Pb(a);if(b!==null&&!e)return[b];var j=Ba(c,e);if(b!==null&&h.inArray(b,j)!==-1)return[b];if(!a)return j;if(typeof a==="function")return h.map(j,function(b){var e=
c.aoData[b];return a(b,e._aData,e.nTr)?b:null});b=Sb(ha(c.aoData,j,"nTr"));if(a.nodeName){if(a._DT_RowIndex!==k)return[a._DT_RowIndex];if(a._DT_CellIndex)return[a._DT_CellIndex.row];b=h(a).closest("*[data-dt-row]");return b.length?[b.data("dt-row")]:[]}if(typeof a==="string"&&a.charAt(0)==="#"){j=c.aIds[a.replace(/^#/,"")];if(j!==k)return[j.idx]}return h(b).filter(a).map(function(){return this._DT_RowIndex}).toArray()},c,e)},1);c.selector.rows=a;c.selector.opts=b;return c});p("rows().nodes()",function(){return this.iterator("row",
function(a,b){return a.aoData[b].nTr||k},1)});p("rows().data()",function(){return this.iterator(!0,"rows",function(a,b){return ha(a.aoData,b,"_aData")},1)});s("rows().cache()","row().cache()",function(a){return this.iterator("row",function(b,c){var d=b.aoData[c];return"search"===a?d._aFilterData:d._aSortData},1)});s("rows().invalidate()","row().invalidate()",function(a){return this.iterator("row",function(b,c){ca(b,c,a)})});s("rows().indexes()","row().index()",function(){return this.iterator("row",
function(a,b){return b},1)});s("rows().ids()","row().id()",function(a){for(var b=[],c=this.context,d=0,e=c.length;d<e;d++)for(var f=0,g=this[d].length;f<g;f++){var h=c[d].rowIdFn(c[d].aoData[this[d][f]]._aData);b.push((!0===a?"#":"")+h)}return new r(c,b)});s("rows().remove()","row().remove()",function(){var a=this;this.iterator("row",function(b,c,d){var e=b.aoData,f=e[c],g,h,i,n,l;e.splice(c,1);g=0;for(h=e.length;g<h;g++)if(i=e[g],l=i.anCells,null!==i.nTr&&(i.nTr._DT_RowIndex=g),null!==l){i=0;for(n=
l.length;i<n;i++)l[i]._DT_CellIndex.row=g}oa(b.aiDisplayMaster,c);oa(b.aiDisplay,c);oa(a[d],c,!1);Sa(b);c=b.rowIdFn(f._aData);c!==k&&delete b.aIds[c]});this.iterator("table",function(a){for(var c=0,d=a.aoData.length;c<d;c++)a.aoData[c].idx=c});return this});p("rows.add()",function(a){var b=this.iterator("table",function(b){var c,f,g,h=[];f=0;for(g=a.length;f<g;f++)c=a[f],c.nodeName&&"TR"===c.nodeName.toUpperCase()?h.push(ma(b,c)[0]):h.push(N(b,c));return h},1),c=this.rows(-1);c.pop();h.merge(c,b);
return c});p("row()",function(a,b){return bb(this.rows(a,b))});p("row().data()",function(a){var b=this.context;if(a===k)return b.length&&this.length?b[0].aoData[this[0]]._aData:k;b[0].aoData[this[0]]._aData=a;ca(b[0],this[0],"data");return this});p("row().node()",function(){var a=this.context;return a.length&&this.length?a[0].aoData[this[0]].nTr||null:null});p("row.add()",function(a){a instanceof h&&a.length&&(a=a[0]);var b=this.iterator("table",function(b){return a.nodeName&&"TR"===a.nodeName.toUpperCase()?
ma(b,a)[0]:N(b,a)});return this.row(b[0])});var cb=function(a,b){var c=a.context;if(c.length&&(c=c[0].aoData[b!==k?b:a[0]])&&c._details)c._details.remove(),c._detailsShow=k,c._details=k},Vb=function(a,b){var c=a.context;if(c.length&&a.length){var d=c[0].aoData[a[0]];if(d._details){(d._detailsShow=b)?d._details.insertAfter(d.nTr):d._details.detach();var e=c[0],f=new r(e),g=e.aoData;f.off("draw.dt.DT_details column-visibility.dt.DT_details destroy.dt.DT_details");0<G(g,"_details").length&&(f.on("draw.dt.DT_details",
function(a,b){e===b&&f.rows({page:"current"}).eq(0).each(function(a){a=g[a];a._detailsShow&&a._details.insertAfter(a.nTr)})}),f.on("column-visibility.dt.DT_details",function(a,b){if(e===b)for(var c,d=aa(b),f=0,h=g.length;f<h;f++)c=g[f],c._details&&c._details.children("td[colspan]").attr("colspan",d)}),f.on("destroy.dt.DT_details",function(a,b){if(e===b)for(var c=0,d=g.length;c<d;c++)g[c]._details&&cb(f,c)}))}}};p("row().child()",function(a,b){var c=this.context;if(a===k)return c.length&&this.length?
c[0].aoData[this[0]]._details:k;if(!0===a)this.child.show();else if(!1===a)cb(this);else if(c.length&&this.length){var d=c[0],c=c[0].aoData[this[0]],e=[],f=function(a,b){if(h.isArray(a)||a instanceof h)for(var c=0,k=a.length;c<k;c++)f(a[c],b);else a.nodeName&&"tr"===a.nodeName.toLowerCase()?e.push(a):(c=h("<tr><td/></tr>").addClass(b),h("td",c).addClass(b).html(a)[0].colSpan=aa(d),e.push(c[0]))};f(a,b);c._details&&c._details.remove();c._details=h(e);c._detailsShow&&c._details.insertAfter(c.nTr)}return this});
p(["row().child.show()","row().child().show()"],function(){Vb(this,!0);return this});p(["row().child.hide()","row().child().hide()"],function(){Vb(this,!1);return this});p(["row().child.remove()","row().child().remove()"],function(){cb(this);return this});p("row().child.isShown()",function(){var a=this.context;return a.length&&this.length?a[0].aoData[this[0]]._detailsShow||!1:!1});var ec=/^(.+):(name|visIdx|visible)$/,Wb=function(a,b,c,d,e){for(var c=[],d=0,f=e.length;d<f;d++)c.push(B(a,e[d],b));
return c};p("columns()",function(a,b){a===k?a="":h.isPlainObject(a)&&(b=a,a="");var b=ab(b),c=this.iterator("table",function(c){var e=a,f=b,g=c.aoColumns,j=G(g,"sName"),i=G(g,"nTh");return $a("column",e,function(a){var b=Pb(a);if(a==="")return W(g.length);if(b!==null)return[b>=0?b:g.length+b];if(typeof a==="function"){var e=Ba(c,f);return h.map(g,function(b,f){return a(f,Wb(c,f,0,0,e),i[f])?f:null})}var k=typeof a==="string"?a.match(ec):"";if(k)switch(k[2]){case "visIdx":case "visible":b=parseInt(k[1],
10);if(b<0){var m=h.map(g,function(a,b){return a.bVisible?b:null});return[m[m.length+b]]}return[Z(c,b)];case "name":return h.map(j,function(a,b){return a===k[1]?b:null});default:return[]}if(a.nodeName&&a._DT_CellIndex)return[a._DT_CellIndex.column];b=h(i).filter(a).map(function(){return h.inArray(this,i)}).toArray();if(b.length||!a.nodeName)return b;b=h(a).closest("*[data-dt-column]");return b.length?[b.data("dt-column")]:[]},c,f)},1);c.selector.cols=a;c.selector.opts=b;return c});s("columns().header()",
"column().header()",function(){return this.iterator("column",function(a,b){return a.aoColumns[b].nTh},1)});s("columns().footer()","column().footer()",function(){return this.iterator("column",function(a,b){return a.aoColumns[b].nTf},1)});s("columns().data()","column().data()",function(){return this.iterator("column-rows",Wb,1)});s("columns().dataSrc()","column().dataSrc()",function(){return this.iterator("column",function(a,b){return a.aoColumns[b].mData},1)});s("columns().cache()","column().cache()",
function(a){return this.iterator("column-rows",function(b,c,d,e,f){return ha(b.aoData,f,"search"===a?"_aFilterData":"_aSortData",c)},1)});s("columns().nodes()","column().nodes()",function(){return this.iterator("column-rows",function(a,b,c,d,e){return ha(a.aoData,e,"anCells",b)},1)});s("columns().visible()","column().visible()",function(a,b){var c=this.iterator("column",function(b,c){if(a===k)return b.aoColumns[c].bVisible;var f=b.aoColumns,g=f[c],j=b.aoData,i,n,l;if(a!==k&&g.bVisible!==a){if(a){var m=
h.inArray(!0,G(f,"bVisible"),c+1);i=0;for(n=j.length;i<n;i++)l=j[i].nTr,f=j[i].anCells,l&&l.insertBefore(f[c],f[m]||null)}else h(G(b.aoData,"anCells",c)).detach();g.bVisible=a;ea(b,b.aoHeader);ea(b,b.aoFooter);wa(b)}});a!==k&&(this.iterator("column",function(c,e){u(c,null,"column-visibility",[c,e,a,b])}),(b===k||b)&&this.columns.adjust());return c});s("columns().indexes()","column().index()",function(a){return this.iterator("column",function(b,c){return"visible"===a?$(b,c):c},1)});p("columns.adjust()",
function(){return this.iterator("table",function(a){Y(a)},1)});p("column.index()",function(a,b){if(0!==this.context.length){var c=this.context[0];if("fromVisible"===a||"toData"===a)return Z(c,b);if("fromData"===a||"toVisible"===a)return $(c,b)}});p("column()",function(a,b){return bb(this.columns(a,b))});p("cells()",function(a,b,c){h.isPlainObject(a)&&(a.row===k?(c=a,a=null):(c=b,b=null));h.isPlainObject(b)&&(c=b,b=null);if(null===b||b===k)return this.iterator("table",function(b){var d=a,e=ab(c),f=
b.aoData,g=Ba(b,e),j=Sb(ha(f,g,"anCells")),i=h([].concat.apply([],j)),l,n=b.aoColumns.length,m,p,r,u,v,s;return $a("cell",d,function(a){var c=typeof a==="function";if(a===null||a===k||c){m=[];p=0;for(r=g.length;p<r;p++){l=g[p];for(u=0;u<n;u++){v={row:l,column:u};if(c){s=f[l];a(v,B(b,l,u),s.anCells?s.anCells[u]:null)&&m.push(v)}else m.push(v)}}return m}if(h.isPlainObject(a))return[a];c=i.filter(a).map(function(a,b){return{row:b._DT_CellIndex.row,column:b._DT_CellIndex.column}}).toArray();if(c.length||
!a.nodeName)return c;s=h(a).closest("*[data-dt-row]");return s.length?[{row:s.data("dt-row"),column:s.data("dt-column")}]:[]},b,e)});var d=this.columns(b,c),e=this.rows(a,c),f,g,j,i,n,l=this.iterator("table",function(a,b){f=[];g=0;for(j=e[b].length;g<j;g++){i=0;for(n=d[b].length;i<n;i++)f.push({row:e[b][g],column:d[b][i]})}return f},1);h.extend(l.selector,{cols:b,rows:a,opts:c});return l});s("cells().nodes()","cell().node()",function(){return this.iterator("cell",function(a,b,c){return(a=a.aoData[b])&&
a.anCells?a.anCells[c]:k},1)});p("cells().data()",function(){return this.iterator("cell",function(a,b,c){return B(a,b,c)},1)});s("cells().cache()","cell().cache()",function(a){a="search"===a?"_aFilterData":"_aSortData";return this.iterator("cell",function(b,c,d){return b.aoData[c][a][d]},1)});s("cells().render()","cell().render()",function(a){return this.iterator("cell",function(b,c,d){return B(b,c,d,a)},1)});s("cells().indexes()","cell().index()",function(){return this.iterator("cell",function(a,
b,c){return{row:b,column:c,columnVisible:$(a,c)}},1)});s("cells().invalidate()","cell().invalidate()",function(a){return this.iterator("cell",function(b,c,d){ca(b,c,a,d)})});p("cell()",function(a,b,c){return bb(this.cells(a,b,c))});p("cell().data()",function(a){var b=this.context,c=this[0];if(a===k)return b.length&&c.length?B(b[0],c[0].row,c[0].column):k;jb(b[0],c[0].row,c[0].column,a);ca(b[0],c[0].row,"data",c[0].column);return this});p("order()",function(a,b){var c=this.context;if(a===k)return 0!==
c.length?c[0].aaSorting:k;"number"===typeof a?a=[[a,b]]:a.length&&!h.isArray(a[0])&&(a=Array.prototype.slice.call(arguments));return this.iterator("table",function(b){b.aaSorting=a.slice()})});p("order.listener()",function(a,b,c){return this.iterator("table",function(d){Ma(d,a,b,c)})});p("order.fixed()",function(a){if(!a){var b=this.context,b=b.length?b[0].aaSortingFixed:k;return h.isArray(b)?{pre:b}:b}return this.iterator("table",function(b){b.aaSortingFixed=h.extend(!0,{},a)})});p(["columns().order()",
"column().order()"],function(a){var b=this;return this.iterator("table",function(c,d){var e=[];h.each(b[d],function(b,c){e.push([c,a])});c.aaSorting=e})});p("search()",function(a,b,c,d){var e=this.context;return a===k?0!==e.length?e[0].oPreviousSearch.sSearch:k:this.iterator("table",function(e){e.oFeatures.bFilter&&fa(e,h.extend({},e.oPreviousSearch,{sSearch:a+"",bRegex:null===b?!1:b,bSmart:null===c?!0:c,bCaseInsensitive:null===d?!0:d}),1)})});s("columns().search()","column().search()",function(a,
b,c,d){return this.iterator("column",function(e,f){var g=e.aoPreSearchCols;if(a===k)return g[f].sSearch;e.oFeatures.bFilter&&(h.extend(g[f],{sSearch:a+"",bRegex:null===b?!1:b,bSmart:null===c?!0:c,bCaseInsensitive:null===d?!0:d}),fa(e,e.oPreviousSearch,1))})});p("state()",function(){return this.context.length?this.context[0].oSavedState:null});p("state.clear()",function(){return this.iterator("table",function(a){a.fnStateSaveCallback.call(a.oInstance,a,{})})});p("state.loaded()",function(){return this.context.length?
this.context[0].oLoadedState:null});p("state.save()",function(){return this.iterator("table",function(a){wa(a)})});m.versionCheck=m.fnVersionCheck=function(a){for(var b=m.version.split("."),a=a.split("."),c,d,e=0,f=a.length;e<f;e++)if(c=parseInt(b[e],10)||0,d=parseInt(a[e],10)||0,c!==d)return c>d;return!0};m.isDataTable=m.fnIsDataTable=function(a){var b=h(a).get(0),c=!1;h.each(m.settings,function(a,e){var f=e.nScrollHead?h("table",e.nScrollHead)[0]:null,g=e.nScrollFoot?h("table",e.nScrollFoot)[0]:
null;if(e.nTable===b||f===b||g===b)c=!0});return c};m.tables=m.fnTables=function(a){var b=!1;h.isPlainObject(a)&&(b=a.api,a=a.visible);var c=h.map(m.settings,function(b){if(!a||a&&h(b.nTable).is(":visible"))return b.nTable});return b?new r(c):c};m.camelToHungarian=K;p("$()",function(a,b){var c=this.rows(b).nodes(),c=h(c);return h([].concat(c.filter(a).toArray(),c.find(a).toArray()))});h.each(["on","one","off"],function(a,b){p(b+"()",function(){var a=Array.prototype.slice.call(arguments);a[0].match(/\.dt\b/)||
(a[0]+=".dt");var d=h(this.tables().nodes());d[b].apply(d,a);return this})});p("clear()",function(){return this.iterator("table",function(a){na(a)})});p("settings()",function(){return new r(this.context,this.context)});p("init()",function(){var a=this.context;return a.length?a[0].oInit:null});p("data()",function(){return this.iterator("table",function(a){return G(a.aoData,"_aData")}).flatten()});p("destroy()",function(a){a=a||!1;return this.iterator("table",function(b){var c=b.nTableWrapper.parentNode,
d=b.oClasses,e=b.nTable,f=b.nTBody,g=b.nTHead,j=b.nTFoot,i=h(e),f=h(f),k=h(b.nTableWrapper),l=h.map(b.aoData,function(a){return a.nTr}),p;b.bDestroying=!0;u(b,"aoDestroyCallback","destroy",[b]);a||(new r(b)).columns().visible(!0);k.unbind(".DT").find(":not(tbody *)").unbind(".DT");h(D).unbind(".DT-"+b.sInstance);e!=g.parentNode&&(i.children("thead").detach(),i.append(g));j&&e!=j.parentNode&&(i.children("tfoot").detach(),i.append(j));b.aaSorting=[];b.aaSortingFixed=[];va(b);h(l).removeClass(b.asStripeClasses.join(" "));
h("th, td",g).removeClass(d.sSortable+" "+d.sSortableAsc+" "+d.sSortableDesc+" "+d.sSortableNone);b.bJUI&&(h("th span."+d.sSortIcon+", td span."+d.sSortIcon,g).detach(),h("th, td",g).each(function(){var a=h("div."+d.sSortJUIWrapper,this);h(this).append(a.contents());a.detach()}));f.children().detach();f.append(l);g=a?"remove":"detach";i[g]();k[g]();!a&&c&&(c.insertBefore(e,b.nTableReinsertBefore),i.css("width",b.sDestroyWidth).removeClass(d.sTable),(p=b.asDestroyStripes.length)&&f.children().each(function(a){h(this).addClass(b.asDestroyStripes[a%
p])}));c=h.inArray(b,m.settings);-1!==c&&m.settings.splice(c,1)})});h.each(["column","row","cell"],function(a,b){p(b+"s().every()",function(a){var d=this.selector.opts,e=this;return this.iterator(b,function(f,g,h,i,n){a.call(e[b](g,"cell"===b?h:d,"cell"===b?d:k),g,h,i,n)})})});p("i18n()",function(a,b,c){var d=this.context[0],a=Q(a)(d.oLanguage);a===k&&(a=b);c!==k&&h.isPlainObject(a)&&(a=a[c]!==k?a[c]:a._);return a.replace("%d",c)});m.version="1.10.12";m.settings=[];m.models={};m.models.oSearch={bCaseInsensitive:!0,
sSearch:"",bRegex:!1,bSmart:!0};m.models.oRow={nTr:null,anCells:null,_aData:[],_aSortData:null,_aFilterData:null,_sFilterRow:null,_sRowStripe:"",src:null,idx:-1};m.models.oColumn={idx:null,aDataSort:null,asSorting:null,bSearchable:null,bSortable:null,bVisible:null,_sManualType:null,_bAttrSrc:!1,fnCreatedCell:null,fnGetData:null,fnSetData:null,mData:null,mRender:null,nTh:null,nTf:null,sClass:null,sContentPadding:null,sDefaultContent:null,sName:null,sSortDataType:"std",sSortingClass:null,sSortingClassJUI:null,
sTitle:null,sType:null,sWidth:null,sWidthOrig:null};m.defaults={aaData:null,aaSorting:[[0,"asc"]],aaSortingFixed:[],ajax:null,aLengthMenu:[10,25,50,100],aoColumns:null,aoColumnDefs:null,aoSearchCols:[],asStripeClasses:null,bAutoWidth:!0,bDeferRender:!1,bDestroy:!1,bFilter:!0,bInfo:!0,bJQueryUI:!1,bLengthChange:!0,bPaginate:!0,bProcessing:!1,bRetrieve:!1,bScrollCollapse:!1,bServerSide:!1,bSort:!0,bSortMulti:!0,bSortCellsTop:!1,bSortClasses:!0,bStateSave:!1,fnCreatedRow:null,fnDrawCallback:null,fnFooterCallback:null,
fnFormatNumber:function(a){return a.toString().replace(/\B(?=(\d{3})+(?!\d))/g,this.oLanguage.sThousands)},fnHeaderCallback:null,fnInfoCallback:null,fnInitComplete:null,fnPreDrawCallback:null,fnRowCallback:null,fnServerData:null,fnServerParams:null,fnStateLoadCallback:function(a){try{return JSON.parse((-1===a.iStateDuration?sessionStorage:localStorage).getItem("DataTables_"+a.sInstance+"_"+location.pathname))}catch(b){}},fnStateLoadParams:null,fnStateLoaded:null,fnStateSaveCallback:function(a,b){try{(-1===
a.iStateDuration?sessionStorage:localStorage).setItem("DataTables_"+a.sInstance+"_"+location.pathname,JSON.stringify(b))}catch(c){}},fnStateSaveParams:null,iStateDuration:7200,iDeferLoading:null,iDisplayLength:10,iDisplayStart:0,iTabIndex:0,oClasses:{},oLanguage:{oAria:{sSortAscending:": activate to sort column ascending",sSortDescending:": activate to sort column descending"},oPaginate:{sFirst:"First",sLast:"Last",sNext:"Next",sPrevious:"Previous"},sEmptyTable:"No data available in table",sInfo:"Showing _START_ to _END_ of _TOTAL_ entries",
sInfoEmpty:"Showing 0 to 0 of 0 entries",sInfoFiltered:"(filtered from _MAX_ total entries)",sInfoPostFix:"",sDecimal:"",sThousands:",",sLengthMenu:"Show _MENU_ entries",sLoadingRecords:"Loading...",sProcessing:"Processing...",sSearch:"Search:",sSearchPlaceholder:"",sUrl:"",sZeroRecords:"No matching records found"},oSearch:h.extend({},m.models.oSearch),sAjaxDataProp:"data",sAjaxSource:null,sDom:"lfrtip",searchDelay:null,sPaginationType:"simple_numbers",sScrollX:"",sScrollXInner:"",sScrollY:"",sServerMethod:"GET",
renderer:null,rowId:"DT_RowId"};X(m.defaults);m.defaults.column={aDataSort:null,iDataSort:-1,asSorting:["asc","desc"],bSearchable:!0,bSortable:!0,bVisible:!0,fnCreatedCell:null,mData:null,mRender:null,sCellType:"td",sClass:"",sContentPadding:"",sDefaultContent:null,sName:"",sSortDataType:"std",sTitle:null,sType:null,sWidth:null};X(m.defaults.column);m.models.oSettings={oFeatures:{bAutoWidth:null,bDeferRender:null,bFilter:null,bInfo:null,bLengthChange:null,bPaginate:null,bProcessing:null,bServerSide:null,
bSort:null,bSortMulti:null,bSortClasses:null,bStateSave:null},oScroll:{bCollapse:null,iBarWidth:0,sX:null,sXInner:null,sY:null},oLanguage:{fnInfoCallback:null},oBrowser:{bScrollOversize:!1,bScrollbarLeft:!1,bBounding:!1,barWidth:0},ajax:null,aanFeatures:[],aoData:[],aiDisplay:[],aiDisplayMaster:[],aIds:{},aoColumns:[],aoHeader:[],aoFooter:[],oPreviousSearch:{},aoPreSearchCols:[],aaSorting:null,aaSortingFixed:[],asStripeClasses:null,asDestroyStripes:[],sDestroyWidth:0,aoRowCallback:[],aoHeaderCallback:[],
aoFooterCallback:[],aoDrawCallback:[],aoRowCreatedCallback:[],aoPreDrawCallback:[],aoInitComplete:[],aoStateSaveParams:[],aoStateLoadParams:[],aoStateLoaded:[],sTableId:"",nTable:null,nTHead:null,nTFoot:null,nTBody:null,nTableWrapper:null,bDeferLoading:!1,bInitialised:!1,aoOpenRows:[],sDom:null,searchDelay:null,sPaginationType:"two_button",iStateDuration:0,aoStateSave:[],aoStateLoad:[],oSavedState:null,oLoadedState:null,sAjaxSource:null,sAjaxDataProp:null,bAjaxDataGet:!0,jqXHR:null,json:k,oAjaxData:k,
fnServerData:null,aoServerParams:[],sServerMethod:null,fnFormatNumber:null,aLengthMenu:null,iDraw:0,bDrawing:!1,iDrawError:-1,_iDisplayLength:10,_iDisplayStart:0,_iRecordsTotal:0,_iRecordsDisplay:0,bJUI:null,oClasses:{},bFiltered:!1,bSorted:!1,bSortCellsTop:null,oInit:null,aoDestroyCallback:[],fnRecordsTotal:function(){return"ssp"==y(this)?1*this._iRecordsTotal:this.aiDisplayMaster.length},fnRecordsDisplay:function(){return"ssp"==y(this)?1*this._iRecordsDisplay:this.aiDisplay.length},fnDisplayEnd:function(){var a=
this._iDisplayLength,b=this._iDisplayStart,c=b+a,d=this.aiDisplay.length,e=this.oFeatures,f=e.bPaginate;return e.bServerSide?!1===f||-1===a?b+d:Math.min(b+a,this._iRecordsDisplay):!f||c>d||-1===a?d:c},oInstance:null,sInstance:null,iTabIndex:0,nScrollHead:null,nScrollFoot:null,aLastSort:[],oPlugins:{},rowIdFn:null,rowId:null};m.ext=v={buttons:{},classes:{},builder:"-source-",errMode:"alert",feature:[],search:[],selector:{cell:[],column:[],row:[]},internal:{},legacy:{ajax:null},pager:{},renderer:{pageButton:{},
header:{}},order:{},type:{detect:[],search:{},order:{}},_unique:0,fnVersionCheck:m.fnVersionCheck,iApiIndex:0,oJUIClasses:{},sVersion:m.version};h.extend(v,{afnFiltering:v.search,aTypes:v.type.detect,ofnSearch:v.type.search,oSort:v.type.order,afnSortData:v.order,aoFeatures:v.feature,oApi:v.internal,oStdClasses:v.classes,oPagination:v.pager});h.extend(m.ext.classes,{sTable:"dataTable",sNoFooter:"no-footer",sPageButton:"paginate_button",sPageButtonActive:"current",sPageButtonDisabled:"disabled",sStripeOdd:"odd",
sStripeEven:"even",sRowEmpty:"dataTables_empty",sWrapper:"dataTables_wrapper",sFilter:"dataTables_filter",sInfo:"dataTables_info",sPaging:"dataTables_paginate paging_",sLength:"dataTables_length",sProcessing:"dataTables_processing",sSortAsc:"sorting_asc",sSortDesc:"sorting_desc",sSortable:"sorting",sSortableAsc:"sorting_asc_disabled",sSortableDesc:"sorting_desc_disabled",sSortableNone:"sorting_disabled",sSortColumn:"sorting_",sFilterInput:"",sLengthSelect:"",sScrollWrapper:"dataTables_scroll",sScrollHead:"dataTables_scrollHead",
sScrollHeadInner:"dataTables_scrollHeadInner",sScrollBody:"dataTables_scrollBody",sScrollFoot:"dataTables_scrollFoot",sScrollFootInner:"dataTables_scrollFootInner",sHeaderTH:"",sFooterTH:"",sSortJUIAsc:"",sSortJUIDesc:"",sSortJUI:"",sSortJUIAscAllowed:"",sSortJUIDescAllowed:"",sSortJUIWrapper:"",sSortIcon:"",sJUIHeader:"",sJUIFooter:""});var Ca="",Ca="",H=Ca+"ui-state-default",ia=Ca+"css_right ui-icon ui-icon-",Xb=Ca+"fg-toolbar ui-toolbar ui-widget-header ui-helper-clearfix";h.extend(m.ext.oJUIClasses,
m.ext.classes,{sPageButton:"fg-button ui-button "+H,sPageButtonActive:"ui-state-disabled",sPageButtonDisabled:"ui-state-disabled",sPaging:"dataTables_paginate fg-buttonset ui-buttonset fg-buttonset-multi ui-buttonset-multi paging_",sSortAsc:H+" sorting_asc",sSortDesc:H+" sorting_desc",sSortable:H+" sorting",sSortableAsc:H+" sorting_asc_disabled",sSortableDesc:H+" sorting_desc_disabled",sSortableNone:H+" sorting_disabled",sSortJUIAsc:ia+"triangle-1-n",sSortJUIDesc:ia+"triangle-1-s",sSortJUI:ia+"carat-2-n-s",
sSortJUIAscAllowed:ia+"carat-1-n",sSortJUIDescAllowed:ia+"carat-1-s",sSortJUIWrapper:"DataTables_sort_wrapper",sSortIcon:"DataTables_sort_icon",sScrollHead:"dataTables_scrollHead "+H,sScrollFoot:"dataTables_scrollFoot "+H,sHeaderTH:H,sFooterTH:H,sJUIHeader:Xb+" ui-corner-tl ui-corner-tr",sJUIFooter:Xb+" ui-corner-bl ui-corner-br"});var Mb=m.ext.pager;h.extend(Mb,{simple:function(){return["previous","next"]},full:function(){return["first","previous","next","last"]},numbers:function(a,b){return[ya(a,
b)]},simple_numbers:function(a,b){return["previous",ya(a,b),"next"]},full_numbers:function(a,b){return["first","previous",ya(a,b),"next","last"]},_numbers:ya,numbers_length:7});h.extend(!0,m.ext.renderer,{pageButton:{_:function(a,b,c,d,e,f){var g=a.oClasses,j=a.oLanguage.oPaginate,i=a.oLanguage.oAria.paginate||{},k,l,m=0,p=function(b,d){var o,r,u,s,v=function(b){Ta(a,b.data.action,true)};o=0;for(r=d.length;o<r;o++){s=d[o];if(h.isArray(s)){u=h("<"+(s.DT_el||"div")+"/>").appendTo(b);p(u,s)}else{k=null;
l="";switch(s){case "ellipsis":b.append('<span class="ellipsis">&#x2026;</span>');break;case "first":k=j.sFirst;l=s+(e>0?"":" "+g.sPageButtonDisabled);break;case "previous":k=j.sPrevious;l=s+(e>0?"":" "+g.sPageButtonDisabled);break;case "next":k=j.sNext;l=s+(e<f-1?"":" "+g.sPageButtonDisabled);break;case "last":k=j.sLast;l=s+(e<f-1?"":" "+g.sPageButtonDisabled);break;default:k=s+1;l=e===s?g.sPageButtonActive:""}if(k!==null){u=h("<a>",{"class":g.sPageButton+" "+l,"aria-controls":a.sTableId,"aria-label":i[s],
"data-dt-idx":m,tabindex:a.iTabIndex,id:c===0&&typeof s==="string"?a.sTableId+"_"+s:null}).html(k).appendTo(b);Wa(u,{action:s},v);m++}}}},r;try{r=h(b).find(I.activeElement).data("dt-idx")}catch(o){}p(h(b).empty(),d);r&&h(b).find("[data-dt-idx="+r+"]").focus()}}});h.extend(m.ext.type.detect,[function(a,b){var c=b.oLanguage.sDecimal;return Za(a,c)?"num"+c:null},function(a){if(a&&!(a instanceof Date)&&(!ac.test(a)||!bc.test(a)))return null;var b=Date.parse(a);return null!==b&&!isNaN(b)||M(a)?"date":
null},function(a,b){var c=b.oLanguage.sDecimal;return Za(a,c,!0)?"num-fmt"+c:null},function(a,b){var c=b.oLanguage.sDecimal;return Rb(a,c)?"html-num"+c:null},function(a,b){var c=b.oLanguage.sDecimal;return Rb(a,c,!0)?"html-num-fmt"+c:null},function(a){return M(a)||"string"===typeof a&&-1!==a.indexOf("<")?"html":null}]);h.extend(m.ext.type.search,{html:function(a){return M(a)?a:"string"===typeof a?a.replace(Ob," ").replace(Aa,""):""},string:function(a){return M(a)?a:"string"===typeof a?a.replace(Ob,
" "):a}});var za=function(a,b,c,d){if(0!==a&&(!a||"-"===a))return-Infinity;b&&(a=Qb(a,b));a.replace&&(c&&(a=a.replace(c,"")),d&&(a=a.replace(d,"")));return 1*a};h.extend(v.type.order,{"date-pre":function(a){return Date.parse(a)||0},"html-pre":function(a){return M(a)?"":a.replace?a.replace(/<.*?>/g,"").toLowerCase():a+""},"string-pre":function(a){return M(a)?"":"string"===typeof a?a.toLowerCase():!a.toString?"":a.toString()},"string-asc":function(a,b){return a<b?-1:a>b?1:0},"string-desc":function(a,
b){return a<b?1:a>b?-1:0}});db("");h.extend(!0,m.ext.renderer,{header:{_:function(a,b,c,d){h(a.nTable).on("order.dt.DT",function(e,f,g,h){if(a===f){e=c.idx;b.removeClass(c.sSortingClass+" "+d.sSortAsc+" "+d.sSortDesc).addClass(h[e]=="asc"?d.sSortAsc:h[e]=="desc"?d.sSortDesc:c.sSortingClass)}})},jqueryui:function(a,b,c,d){h("<div/>").addClass(d.sSortJUIWrapper).append(b.contents()).append(h("<span/>").addClass(d.sSortIcon+" "+c.sSortingClassJUI)).appendTo(b);h(a.nTable).on("order.dt.DT",function(e,
f,g,h){if(a===f){e=c.idx;b.removeClass(d.sSortAsc+" "+d.sSortDesc).addClass(h[e]=="asc"?d.sSortAsc:h[e]=="desc"?d.sSortDesc:c.sSortingClass);b.find("span."+d.sSortIcon).removeClass(d.sSortJUIAsc+" "+d.sSortJUIDesc+" "+d.sSortJUI+" "+d.sSortJUIAscAllowed+" "+d.sSortJUIDescAllowed).addClass(h[e]=="asc"?d.sSortJUIAsc:h[e]=="desc"?d.sSortJUIDesc:c.sSortingClassJUI)}})}}});var Yb=function(a){return"string"===typeof a?a.replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/"/g,"&quot;"):a};m.render={number:function(a,
b,c,d,e){return{display:function(f){if("number"!==typeof f&&"string"!==typeof f)return f;var g=0>f?"-":"",h=parseFloat(f);if(isNaN(h))return Yb(f);f=Math.abs(h);h=parseInt(f,10);f=c?b+(f-h).toFixed(c).substring(2):"";return g+(d||"")+h.toString().replace(/\B(?=(\d{3})+(?!\d))/g,a)+f+(e||"")}}},text:function(){return{display:Yb}}};h.extend(m.ext.internal,{_fnExternApiFunc:Nb,_fnBuildAjax:ra,_fnAjaxUpdate:lb,_fnAjaxParameters:ub,_fnAjaxUpdateDraw:vb,_fnAjaxDataSrc:sa,_fnAddColumn:Ea,_fnColumnOptions:ja,
_fnAdjustColumnSizing:Y,_fnVisibleToColumnIndex:Z,_fnColumnIndexToVisible:$,_fnVisbleColumns:aa,_fnGetColumns:la,_fnColumnTypes:Ga,_fnApplyColumnDefs:ib,_fnHungarianMap:X,_fnCamelToHungarian:K,_fnLanguageCompat:Da,_fnBrowserDetect:gb,_fnAddData:N,_fnAddTr:ma,_fnNodeToDataIndex:function(a,b){return b._DT_RowIndex!==k?b._DT_RowIndex:null},_fnNodeToColumnIndex:function(a,b,c){return h.inArray(c,a.aoData[b].anCells)},_fnGetCellData:B,_fnSetCellData:jb,_fnSplitObjNotation:Ja,_fnGetObjectDataFn:Q,_fnSetObjectDataFn:R,
_fnGetDataMaster:Ka,_fnClearTable:na,_fnDeleteIndex:oa,_fnInvalidate:ca,_fnGetRowElements:Ia,_fnCreateTr:Ha,_fnBuildHead:kb,_fnDrawHead:ea,_fnDraw:O,_fnReDraw:T,_fnAddOptionsHtml:nb,_fnDetectHeader:da,_fnGetUniqueThs:qa,_fnFeatureHtmlFilter:pb,_fnFilterComplete:fa,_fnFilterCustom:yb,_fnFilterColumn:xb,_fnFilter:wb,_fnFilterCreateSearch:Pa,_fnEscapeRegex:Qa,_fnFilterData:zb,_fnFeatureHtmlInfo:sb,_fnUpdateInfo:Cb,_fnInfoMacros:Db,_fnInitialise:ga,_fnInitComplete:ta,_fnLengthChange:Ra,_fnFeatureHtmlLength:ob,
_fnFeatureHtmlPaginate:tb,_fnPageChange:Ta,_fnFeatureHtmlProcessing:qb,_fnProcessingDisplay:C,_fnFeatureHtmlTable:rb,_fnScrollDraw:ka,_fnApplyToChildren:J,_fnCalculateColumnWidths:Fa,_fnThrottle:Oa,_fnConvertToWidth:Fb,_fnGetWidestNode:Gb,_fnGetMaxLenString:Hb,_fnStringToCss:x,_fnSortFlatten:V,_fnSort:mb,_fnSortAria:Jb,_fnSortListener:Va,_fnSortAttachListener:Ma,_fnSortingClasses:va,_fnSortData:Ib,_fnSaveState:wa,_fnLoadState:Kb,_fnSettingsFromNode:xa,_fnLog:L,_fnMap:E,_fnBindAction:Wa,_fnCallbackReg:z,
_fnCallbackFire:u,_fnLengthOverflow:Sa,_fnRenderer:Na,_fnDataSource:y,_fnRowAttributes:La,_fnCalculateEnd:function(){}});h.fn.dataTable=m;m.$=h;h.fn.dataTableSettings=m.settings;h.fn.dataTableExt=m.ext;h.fn.DataTable=function(a){return h(this).dataTable(a).api()};h.each(m,function(a,b){h.fn.DataTable[a]=b});return h.fn.dataTable});

$.extend(!0,$.fn.dataTable.defaults,{sDom:"<'row'<'col-xs-6'l><'col-xs-6'f>r>t<'row'<'col-xs-6'i><'col-xs-6'p>>",oLanguage:{sLengthMenu:"Display _MENU_ records"}}),$.extend($.fn.dataTableExt.oStdClasses,{sWrapper:"dataTables_wrapper form-inline",sFilterInput:"form-control input-sm",sLengthSelect:"form-control input-sm"}),$.fn.dataTable.Api?($.fn.dataTable.defaults.renderer="bootstrap",$.fn.dataTable.ext.renderer.pageButton.bootstrap=function(a,b,c,d,e,f){var g,h,i=new $.fn.dataTable.Api(a),j=a.oClasses,k=a.oLanguage.oPaginate,l=function(b,d){var m,n,o,p,q=function(a){return a.preventDefault(),$(a.target).parent().hasClass("disabled")?!1:void("ellipsis"!==a.data.action&&i.page(a.data.action).draw(!1))};for(m=0,n=d.length;n>m;m++)if(p=d[m],$.isArray(p))l(b,p);else{switch(g="",h="",p){case"ellipsis":g="&hellip;",h="disabled";break;case"first":g=k.sFirst,h=p+(e>0?"":" disabled");break;case"previous":g=k.sPrevious,h=p+(e>0?"":" disabled");break;case"next":g=k.sNext,h=p+(f-1>e?"":" disabled");break;case"last":g=k.sLast,h=p+(f-1>e?"":" disabled");break;default:g=p+1,h=e===p?"active":""}g&&(o=$("<li>",{"class":j.sPageButton+" "+h,"aria-controls":a.sTableId,tabindex:a.iTabIndex,id:0===c&&"string"==typeof p?a.sTableId+"_"+p:null}).append($("<a>",{href:"#"}).html(g)).appendTo(b),a.oApi._fnBindAction(o,{action:p},q))}};l($(b).empty().html('<ul class="pagination"/>').children("ul"),d)}):($.fn.dataTable.defaults.sPaginationType="bootstrap",$.fn.dataTableExt.oApi.fnPagingInfo=function(a){return{iStart:a._iDisplayStart,iEnd:a.fnDisplayEnd(),iLength:a._iDisplayLength,iTotal:a.fnRecordsTotal(),iFilteredTotal:a.fnRecordsDisplay(),iPage:-1===a._iDisplayLength?0:Math.ceil(a._iDisplayStart/a._iDisplayLength),iTotalPages:-1===a._iDisplayLength?0:Math.ceil(a.fnRecordsDisplay()/a._iDisplayLength)}},$.extend($.fn.dataTableExt.oPagination,{bootstrap:{fnInit:function(a,b,c){var d=(a.oLanguage.oPaginate,function(b){alert(1),b.preventDefault(),a.oApi._fnPageChange(a,b.data.action)&&c(a)});$(b).append('<ul class="pagination"><li class="prev disabled"><a href="#"><i class="fa fa-angle-double-left"></i></a></li><li class="prev disabled"><a href="#"><i class="fa fa-angle-left"></i></a></li><li class="next disabled"><a href="#"><i class="fa fa-angle-right"></i></a></li><li class="next disabled"><a href="#"><i class="fa fa-angle-double-right"></i></a></li></ul>');var e=$("a",b);$(e[0]).bind("click.DT",{action:"first"},d),$(e[1]).bind("click.DT",{action:"previous"},d),$(e[2]).bind("click.DT",{action:"next"},d),$(e[3]).bind("click.DT",{action:"last"},d)},fnUpdate:function(a,b){var c,d,e,f,g,h,i=5,j=a.oInstance.fnPagingInfo(),k=a.aanFeatures.p,l=Math.floor(i/2);for(j.iTotalPages<i?(g=1,h=j.iTotalPages):j.iPage<=l?(g=1,h=i):j.iPage>=j.iTotalPages-l?(g=j.iTotalPages-i+1,h=j.iTotalPages):(g=j.iPage-l+1,h=g+i-1),c=0,d=k.length;d>c;c++){for($("li:gt(0)",k[c]).filter(":not(.next,.prev)").remove(),e=g;h>=e;e++)f=e==j.iPage+1?'class="active"':"",$("<li "+f+'><a href="#">'+e+"</a></li>").insertBefore($("li.next:eq(0)",k[c])[0]).bind("click",function(c){c.preventDefault(),a._iDisplayStart=(parseInt($("a",this).text(),10)-1)*j.iLength,b(a)});0===j.iPage?$("li.prev",k[c]).addClass("disabled"):$("li.prev",k[c]).removeClass("disabled"),j.iPage===j.iTotalPages-1||0===j.iTotalPages?$("li.next",k[c]).addClass("disabled"):$("li.next",k[c]).removeClass("disabled")}}}})),$.fn.DataTable.TableTools&&($.extend(!0,$.fn.DataTable.TableTools.classes,{container:"DTTT btn-group",buttons:{normal:"btn btn-default",disabled:"disabled"},collection:{container:"DTTT_dropdown dropdown-menu",buttons:{normal:"",disabled:"disabled"}},print:{info:"DTTT_print_info modal"},select:{row:"active"}}),$.extend(!0,$.fn.DataTable.TableTools.DEFAULTS.oTags,{collection:{container:"ul",button:"li",liner:"a"}}));
/*! Buttons for DataTables 1.0.4-dev
 * ©2015 SpryMedia Ltd - datatables.net/license
 */
!function(a,b,c){var d=function(d,e){"use strict";var f=0,g=0,h=e.ext.buttons,i=function(a,b){b===!0&&(b={}),d.isArray(b)&&(b={buttons:b}),this.c=d.extend(!0,{},i.defaults,b),b.buttons&&(this.c.buttons=b.buttons),this.s={dt:new e.Api(a),buttons:[],subButtons:[],listenKeys:"",namespace:"dtb"+f++},this.dom={container:d("<"+this.c.dom.container.tag+"/>").addClass(this.c.dom.container.className)},this._constructor()};d.extend(i.prototype,{action:function(a,b){var d=this._indexToButton(a).conf;return b===c?d.action:(d.action=b,this)},active:function(a,b){var d=this._indexToButton(a);return d.node.toggleClass(this.c.dom.button.active,b===c?!0:b),this},add:function(a,b){if("string"==typeof a&&-1!==a.indexOf("-")){var c=a.split("-");this.c.buttons[1*c[0]].buttons.splice(1*c[1],0,b)}else this.c.buttons.splice(1*a,0,b);return this.dom.container.empty(),this._buildButtons(this.c.buttons),this},container:function(){return this.dom.container},disable:function(a){var b=this._indexToButton(a);return b.node.addClass(this.c.dom.button.disabled),this},destroy:function(){d("body").off("keyup."+this.s.namespace);var a,b,c,e,f=this.s.buttons,g=this.s.subButtons;for(a=0,b=f.length;b>a;a++)for(this.removePrep(a),c=0,e=g[a].length;e>c;c++)this.removePrep(a+"-"+c);this.removeCommit(),this.dom.container.remove();var h=this.s.dt.settings()[0];for(a=0,b=h.length;b>a;a++)if(h.inst===this){h.splice(a,1);break}return this},enable:function(a,b){if(b===!1)return this.disable(a);var c=this._indexToButton(a);return c.node.removeClass(this.c.dom.button.disabled),this},name:function(){return this.c.name},node:function(a){var b=this._indexToButton(a);return b.node},removeCommit:function(){var a,b,c,d=this.s.buttons,e=this.s.subButtons;for(a=d.length-1;a>=0;a--)null===d[a]&&(d.splice(a,1),e.splice(a,1),this.c.buttons.splice(a,1));for(a=0,b=e.length;b>a;a++)for(c=e[a].length-1;c>=0;c--)null===e[a][c]&&(e[a].splice(c,1),this.c.buttons[a].buttons.splice(c,1));return this},removePrep:function(a){var b,c=this.s.dt;if("number"==typeof a||-1===a.indexOf("-"))b=this.s.buttons[1*a],b.conf.destroy&&b.conf.destroy.call(c.button(a),c,b,b.conf),b.node.remove(),this._removeKey(b.conf),this.s.buttons[1*a]=null;else{var d=a.split("-");b=this.s.subButtons[1*d[0]][1*d[1]],b.conf.destroy&&b.conf.destroy.call(c.button(a),c,b,b.conf),b.node.remove(),this._removeKey(b.conf),this.s.subButtons[1*d[0]][1*d[1]]=null}return this},text:function(a,b){var d=this._indexToButton(a),e=this.c.dom.buttonLiner.tag,f=this.s.dt,g=function(a){return"function"==typeof a?a(f,d.node,d.conf):a};return b===c?g(d.conf.text):(d.conf.text=b,e?d.node.children(e).html(g(b)):d.node.html(g(b)),this)},toIndex:function(a){var b,c,d,e,f=this.s.buttons,g=this.s.subButtons;for(b=0,c=f.length;c>b;b++)if(f[b].node[0]===a)return b+"";for(b=0,c=g.length;c>b;b++)for(d=0,e=g[b].length;e>d;d++)if(g[b][d].node[0]===a)return b+"-"+d},_constructor:function(){var a=this,c=this.s.dt,e=c.settings()[0];e._buttons||(e._buttons=[]),e._buttons.push({inst:this,name:this.c.name}),this._buildButtons(this.c.buttons),c.on("destroy",function(){a.destroy()}),d("body").on("keyup."+this.s.namespace,function(c){if(!b.activeElement||b.activeElement===b.body){var d=String.fromCharCode(c.keyCode).toLowerCase();-1!==a.s.listenKeys.toLowerCase().indexOf(d)&&a._keypress(d,c)}})},_addKey:function(a){a.key&&(this.s.listenKeys+=d.isPlainObject(a.key)?a.key.key:a.key)},_buildButtons:function(a,b,e){var f=this.s.dt;b||(b=this.dom.container,this.s.buttons=[],this.s.subButtons=[]);for(var g=0,h=a.length;h>g;g++){var i=this._resolveExtends(a[g]);if(i)if(d.isArray(i))this._buildButtons(i,b,e);else{var j=this._buildButton(i,e!==c?!0:!1);if(j){var k=j.node;if(b.append(j.inserter),e===c?(this.s.buttons.push({node:k,conf:i,inserter:j.inserter}),this.s.subButtons.push([])):this.s.subButtons[e].push({node:k,conf:i,inserter:j.inserter}),i.buttons){var l=this.c.dom.collection;i._collection=d("<"+l.tag+"/>").addClass(l.className),this._buildButtons(i.buttons,i._collection,g)}i.init&&i.init.call(f.button(k),f,k,i)}}}},_buildButton:function(a,b){var c=this.c.dom.button,e=this.c.dom.buttonLiner,f=this.c.dom.collection,h=this.s.dt,i=function(b){return"function"==typeof b?b(h,j,a):b};if(b&&f.button&&(c=f.button),b&&f.buttonLiner&&(e=f.buttonLiner),a.available&&!a.available(h,a))return!1;var j=d("<"+c.tag+"/>").addClass(c.className).attr("tabindex",this.s.dt.settings()[0].iTabIndex).attr("aria-controls",this.s.dt.table().node().id).on("click.dtb",function(b){b.preventDefault(),!j.hasClass(c.disabled)&&a.action&&a.action.call(h.button(j),b,h,j,a),j.blur()}).on("keyup.dtb",function(b){13===b.keyCode&&!j.hasClass(c.disabled)&&a.action&&a.action.call(h.button(j),b,h,j,a)});e.tag?j.append(d("<"+e.tag+"/>").html(i(a.text)).addClass(e.className)):j.html(i(a.text)),a.enabled===!1&&j.addClass(c.disabled),a.className&&j.addClass(a.className),a.namespace||(a.namespace=".dt-button-"+g++);var k,l=this.c.dom.buttonContainer;return k=l?d("<"+l.tag+"/>").addClass(l.className).append(j):j,this._addKey(a),{node:j,inserter:k}},_indexToButton:function(a){if("number"==typeof a||-1===a.indexOf("-"))return this.s.buttons[1*a];var b=a.split("-");return this.s.subButtons[1*b[0]][1*b[1]]},_keypress:function(a,b){var c,e,f,g,h=this.s.buttons,i=this.s.subButtons,j=function(c,e){if(c.key)if(c.key===a)e.click();else if(d.isPlainObject(c.key)){if(c.key.key!==a)return;if(c.key.shiftKey&&!b.shiftKey)return;if(c.key.altKey&&!b.altKey)return;if(c.key.ctrlKey&&!b.ctrlKey)return;if(c.key.metaKey&&!b.metaKey)return;e.click()}};for(c=0,e=h.length;e>c;c++)j(h[c].conf,h[c].node);for(c=0,e=i.length;e>c;c++)for(f=0,g=i[c].length;g>f;f++)j(i[c][f].conf,i[c][f].node)},_removeKey:function(a){if(a.key){var b=d.isPlainObject(a.key)?a.key.key:a.key,c=this.s.listenKeys.split(""),e=d.inArray(b,c);c.splice(e,1),this.s.listenKeys=c.join("")}},_resolveExtends:function(a){var b,e,f=this.s.dt,g=function(b){for(var e=0;!d.isPlainObject(b)&&!d.isArray(b);){if(b===c)return;if("function"==typeof b){if(b=b(f,a),!b)return!1}else if("string"==typeof b){if(!h[b])throw"Unknown button type: "+b;b=h[b]}if(e++,e>30)throw"Buttons: Too many iterations"}return d.isArray(b)?b:d.extend({},b)};for(a=g(a);a&&a.extend;){if(!h[a.extend])throw"Cannot extend unknown button type: "+a.extend;var i=g(h[a.extend]);if(d.isArray(i))return i;var j=i.className;a=d.extend({},i,a),j&&a.className!==j&&(a.className=j+" "+a.className);var k=a.postfixButtons;if(k){for(a.buttons||(a.buttons=[]),b=0,e=k.length;e>b;b++)a.buttons.push(k[b]);a.postfixButtons=null}var l=a.prefixButtons;if(l){for(a.buttons||(a.buttons=[]),b=0,e=l.length;e>b;b++)a.buttons.splice(b,0,l[b]);a.prefixButtons=null}a.extend=i.extend}return a}}),i.background=function(a,b,e){e===c&&(e=400),a?d("<div/>").addClass(b).css("display","none").appendTo("body").fadeIn(e):d("body > div."+b).fadeOut(e,function(){d(this).remove()})},i.instanceSelector=function(a,b){if(!a)return d.map(b,function(a){return a.inst});var c=[],e=d.map(b,function(a){return a.name}),f=function(a){if(d.isArray(a))for(var g=0,h=a.length;h>g;g++)f(a[g]);else if("string"==typeof a)if(-1!==a.indexOf(","))f(a.split(","));else{var i=d.inArray(d.trim(a),e);-1!==i&&c.push(b[i].inst)}else"number"==typeof a&&c.push(b[a].inst)};return f(a),c},i.buttonSelector=function(a,b){for(var e=[],f=function(a,b){var g,h,i=[];d.each(b.s.buttons,function(a,b){null!==b&&i.push({node:b.node[0],name:b.name})}),d.each(b.s.subButtons,function(a,b){d.each(b,function(a,b){null!==b&&i.push({node:b.node[0],name:b.name})})});var j=d.map(i,function(a){return a.node});if(d.isArray(a)||a instanceof d)for(g=0,h=a.length;h>g;g++)f(a[g],b);else if(null===a||a===c||"*"===a)for(g=0,h=i.length;h>g;g++)e.push({inst:b,idx:b.toIndex(i[g].node)});else if("number"==typeof a)e.push({inst:b,idx:a});else if("string"==typeof a)if(-1!==a.indexOf(",")){var k=a.split(",");for(g=0,h=k.length;h>g;g++)f(d.trim(k[g]),b)}else if(a.match(/^\d+(\-\d+)?$/))e.push({inst:b,idx:a});else if(-1!==a.indexOf(":name")){var l=a.replace(":name","");for(g=0,h=i.length;h>g;g++)i[g].name===l&&e.push({inst:b,idx:b.toIndex(i[g].node)})}else d(j).filter(a).each(function(){e.push({inst:b,idx:b.toIndex(this)})});else if("object"==typeof a&&a.nodeName){var m=d.inArray(a,j);-1!==m&&e.push({inst:b,idx:b.toIndex(j[m])})}},g=0,h=a.length;h>g;g++){var i=a[g];f(b,i)}return e},i.defaults={buttons:["copy","excel","csv","pdf","print"],name:"main",tabIndex:0,dom:{container:{tag:"div",className:"dt-buttons"},collection:{tag:"div",className:"dt-button-collection"},button:{tag:"a",className:"dt-button",active:"active",disabled:"disabled"},buttonLiner:{tag:"span",className:""}}},i.version="1.0.4-dev",d.extend(h,{collection:{text:function(a,b,c){return a.i18n("buttons.collection","Collection")},className:"buttons-collection",action:function(c,e,f,g){var h=f,j=h.offset(),k=d(e.table().container());if(g._collection.addClass(g.collectionLayout).css("display","none").appendTo("body").fadeIn(g.fade),"absolute"===g._collection.css("position")){g._collection.css({top:j.top+h.outerHeight(),left:j.left});var l=j.left+g._collection.outerWidth(),m=k.offset().left+k.width();l>m&&g._collection.css("left",j.left-(l-m))}else{var n=g._collection.height()/2;n>d(a).height()/2&&(n=d(a).height()/2),g._collection.css("marginTop",-1*n)}g.background&&i.background(!0,g.backgroundClassName,g.fade),setTimeout(function(){d(b).on("click.dtb-collection",function(a){d(a.target).parents().andSelf().filter(g._collection).length||(g._collection.fadeOut(g.fade,function(){g._collection.detach()}),i.background(!1,g.backgroundClassName,g.fade),d(b).off("click.dtb-collection"))})},10)},background:!0,collectionLayout:"",backgroundClassName:"dt-button-background",fade:400},copy:function(a,b){return b.preferHtml&&h.copyHtml5?"copyHtml5":h.copyFlash&&h.copyFlash.available(a,b)?"copyFlash":h.copyHtml5?"copyHtml5":void 0},csv:function(a,b){return h.csvHtml5&&h.csvHtml5.available(a,b)?"csvHtml5":h.csvFlash&&h.csvFlash.available(a,b)?"csvFlash":void 0},excel:function(a,b){return h.excelHtml5&&h.excelHtml5.available(a,b)?"excelHtml5":h.excelFlash&&h.excelFlash.available(a,b)?"excelFlash":void 0},pdf:function(a,b){return h.pdfHtml5&&h.pdfHtml5.available(a,b)?"pdfHtml5":h.pdfFlash&&h.pdfFlash.available(a,b)?"pdfFlash":void 0}}),e.Api.register("buttons()",function(a,b){return b===c&&(b=a,a=c),this.iterator(!0,"table",function(c){return c._buttons?i.buttonSelector(i.instanceSelector(a,c._buttons),b):void 0},!0)}),e.Api.register("button()",function(a,b){var c=this.buttons(a,b);return c.length>1&&c.splice(1,c.length),c}),e.Api.register(["buttons().active()","button().active()"],function(a){return this.each(function(b){b.inst.active(b.idx,a)})}),e.Api.registerPlural("buttons().action()","button().action()",function(a){return a===c?this.map(function(a){return a.inst.action(a.idx)}):this.each(function(b){b.inst.action(b.idx,a)})}),e.Api.register(["buttons().enable()","button().enable()"],function(a){return this.each(function(b){b.inst.enable(b.idx,a)})}),e.Api.register(["buttons().disable()","button().disable()"],function(){return this.each(function(a){a.inst.disable(a.idx)})}),e.Api.registerPlural("buttons().nodes()","button().node()",function(){var a=d();return d(this.each(function(b){a=a.add(b.inst.node(b.idx))})),a}),e.Api.registerPlural("buttons().text()","button().text()",function(a){return a===c?this.map(function(a){return a.inst.text(a.idx)}):this.each(function(b){b.inst.text(b.idx,a)})}),e.Api.registerPlural("buttons().trigger()","button().trigger()",function(){return this.each(function(a){a.inst.node(a.idx).trigger("click")})}),e.Api.registerPlural("buttons().containers()","buttons().container()",function(){var a=d();return d(this.each(function(b){a=a.add(b.inst.container())})),a}),e.Api.register("button().add()",function(a,b){return 1===this.length&&this[0].inst.add(a,b),this.button(a)}),e.Api.register("buttons().destroy()",function(a){return this.pluck("inst").unique().each(function(a){a.destroy()}),this}),e.Api.registerPlural("buttons().remove()","buttons().remove()",function(){return this.each(function(a){a.inst.removePrep(a.idx)}),this.pluck("inst").unique().each(function(a){a.removeCommit()}),this});var j;e.Api.register("buttons.info()",function(a,b,e){var f=this;return a===!1?(d("#datatables_buttons_info").fadeOut(function(){d(this).remove()}),clearTimeout(j),j=null,this):(j&&clearTimeout(j),d("#datatables_buttons_info").length&&d("#datatables_buttons_info").remove(),a=a?"<h2>"+a+"</h2>":"",d('<div id="datatables_buttons_info" class="dt-button-info"/>').html(a).append(d("<div/>")["string"==typeof b?"html":"append"](b)).css("display","none").appendTo("body").fadeIn(),e!==c&&0!==e&&(j=setTimeout(function(){f.buttons.info(!1)},e)),this)}),e.Api.register("buttons.exportData()",function(a){return this.context.length?k(new e.Api(this.context[0]),a):void 0});var k=function(a,b){for(var c=d.extend(!0,{},{rows:null,columns:"",modifier:{search:"applied",order:"applied"},orthogonal:"display",stripHtml:!0,stripNewlines:!0,trim:!0},b),e=function(a){return"string"!=typeof a?a:(c.stripHtml&&(a=a.replace(/<.*?>/g,"")),c.trim&&(a=a.replace(/^\s+|\s+$/g,"")),c.stripNewlines&&(a=a.replace(/\n/g," ")),a)},f=a.columns(c.columns).indexes().map(function(b,c){return e(a.column(b).header().innerHTML)}).toArray(),g=a.table().footer()?a.columns(c.columns).indexes().map(function(b,c){var d=a.column(b).footer();return d?e(d.innerHTML):""}).toArray():null,h=a.rows(c.rows,c.modifier).indexes().toArray(),i=a.cells(h,c.columns).render(c.orthogonal).toArray(),j=f.length,k=j>0?i.length/j:0,l=new Array(k),m=0,n=0,o=k;o>n;n++){for(var p=new Array(j),q=0;j>q;q++)p[q]=e(i[m]),m++;l[n]=p}return{header:f,footer:g,body:l}};return d.fn.dataTable.Buttons=i,d.fn.DataTable.Buttons=i,d(b).on("init.dt.dtb",function(a,b,c){if("dt"===a.namespace){var d=b.oInit.buttons||e.defaults.buttons;d&&!b._buttons&&new i(b,d).container()}}),e.ext.feature.push({fnInit:function(a){var b=new e.Api(a),c=b.init().buttons;return new i(b,c).container()},cFeature:"B"}),i};"function"==typeof define&&define.amd?define(["jquery","datatables"],d):"object"==typeof exports?d(require("jquery"),require("datatables")):jQuery&&!jQuery.fn.dataTable.Buttons&&d(jQuery,jQuery.fn.dataTable)}(window,document);
/*!
 * Column visibility buttons for Buttons and DataTables.
 * 2015 SpryMedia Ltd - datatables.net/license
 */
!function(a,b){"use strict";a.extend(b.ext.buttons,{colvis:function(a,b){return{extend:"collection",text:function(a){return a.i18n("buttons.colvis","Column visibility")},className:"buttons-colvis",buttons:[{extend:"columnsToggle",columns:b.columns}]}},columnsToggle:function(a,b){var c=a.columns(b.columns).indexes().map(function(a){return{extend:"columnToggle",columns:a}}).toArray();return c},columnToggle:function(a,b){return{extend:"columnVisibility",columns:b.columns}},columnsVisibility:function(a,b){var c=a.columns(b.columns).indexes().map(function(a){return{extend:"columnVisibility",columns:a,visibility:b.visibility}}).toArray();return c},columnVisibility:{columns:null,text:function(a,b,c){return c._columnText(a,c.columns)},className:"buttons-columnVisibility",action:function(a,b,c,d){var e=b.column(d.columns);e.visible(void 0!==d.visibility?d.visibility:!e.visible())},init:function(a,b,c){var d=this,e=a.column(c.columns);a.on("column-visibility.dt"+c.namespace,function(a,b,e,f){e===c.columns&&d.active(f)}).on("column-reorder.dt"+c.namespace,function(e,f,g){var h=a.column(c.columns);b.text(c._columnText(a,c.columns)),d.active(h.visible())}),this.active(e.visible())},destroy:function(a,b,c){a.off("column-visibility.dt"+c.namespace).off("column-reorder.dt"+c.namespace)},_columnText:function(a,b){var c=a.column(b).index();return a.settings()[0].aoColumns[c].sTitle.replace(/\n/g," ").replace(/<.*?>/g,"").replace(/^\s+|\s+$/g,"")}},colvisRestore:{className:"buttons-colvisRestore",text:function(a){return a.i18n("buttons.colvisRestore","Restore visibility")},init:function(a,b,c){c._visOriginal=a.columns().indexes().map(function(b){return a.column(b).visible()}).toArray()},action:function(a,b,c,d){b.columns().every(function(a){this.visible(d._visOriginal[a])})}},colvisGroup:{className:"buttons-colvisGroup",action:function(a,b,c,d){b.columns(d.show).visible(!0),b.columns(d.hide).visible(!1)},show:[],hide:[]}})}(jQuery,jQuery.fn.dataTable);
/*! Select for DataTables 1.0.1
 * 2015 SpryMedia Ltd - datatables.net/license/mit
 */
!function(a,b,c){var d=function(a,d){"use strict";function e(a,b,c){var d,e,f,g=function(b,c){if(b>c){var d=c;c=b,b=d}var e=!1;return a.columns(":visible").indexes().filter(function(a){return a===b&&(e=!0),a===c?(e=!1,!0):e})},h=function(b,c){var d=a.rows({search:"applied"}).indexes();if(d.indexOf(b)>d.indexOf(c)){var e=c;c=b,b=e}var f=!1;return d.filter(function(a){return a===b&&(f=!0),a===c?(f=!1,!0):f})};a.cells({selected:!0}).any()||c?(e=g(c.column,b.column),f=h(c.row,b.row)):(e=g(0,b.column),f=h(0,b.row)),d=a.cells(f,e).flatten(),a.cells(b,{selected:!0}).any()?a.cells(d).deselect():a.cells(d).select()}function f(b){var c=b.settings()[0],d=c._select.selector;a(b.table().body()).off("mousedown.dtSelect",d).off("mouseup.dtSelect",d).off("click.dtSelect",d),a("body").off("click.dtSelect")}function g(b){var c=a(b.table().body()),d=b.settings()[0],e=d._select.selector;c.on("mousedown.dtSelect",e,function(a){a.shiftKey&&c.css("-moz-user-select","none").one("selectstart.dtSelect",e,function(){return!1})}).on("mouseup.dtSelect",e,function(a){c.css("-moz-user-select","")}).on("click.dtSelect",e,function(d){var e,f=b.select.items(),g=b.cell(this).index(),h=b.settings()[0];a(d.target).closest("tbody")[0]==c[0]&&b.cell(d.target).any()&&("row"===f?(e=g.row,m(d,b,h,"row",e)):"column"===f?(e=b.cell(d.target).index().column,m(d,b,h,"column",e)):"cell"===f&&(e=b.cell(d.target).index(),m(d,b,h,"cell",e)),h._select_lastCell=g)}),a("body").on("click.dtSelect",function(c){if(d._select.blurable){if(a(c.target).parents().filter(b.table().container()).length)return;if(a(c.target).parents("div.DTE").length)return;l(d,!0)}})}function h(b,c,d,e){(!e||b.flatten().length)&&(d.unshift(b),a(b.table().node()).triggerHandler(c+".dt",d))}function i(b){var c=b.settings()[0];if(c._select.info&&c.aanFeatures.i){var d=a('<span class="select-info"/>'),e=function(c,e){d.append(a('<span class="select-item"/>').append(b.i18n("select."+c+"s",{_:"%d "+c+"s selected",0:"",1:"1 "+c+" selected"},e)))};e("row",b.rows({selected:!0}).flatten().length),e("column",b.columns({selected:!0}).flatten().length),e("cell",b.cells({selected:!0}).flatten().length),a.each(c.aanFeatures.i,function(b,c){c=a(c);var e=c.children("span.select-info");e.length&&e.remove(),""!==d.text()&&c.append(d)})}}function j(b){var e=new d.Api(b);b.aoRowCreatedCallback.push({fn:function(c,d,e){var f,g,h=b.aoData[e];for(h._select_selected&&a(c).addClass("selected"),f=0,g=b.aoColumns.length;g>f;f++)(b.aoColumns[f]._select_selected||h._selected_cells&&h._selected_cells[f])&&a(h.anCells[f]).addClass("selected")},sName:"select-deferRender"}),e.on("preXhr.dt.dtSelect",function(){var a=e.rows({selected:!0}).ids(!0).filter(function(a){return a!==c}),b=e.cells({selected:!0}).eq(0).map(function(a){var b=e.row(a.row).id(!0);return b?{row:b,column:a.column}:c}).filter(function(a){return a!==c});e.one("draw.dt.dtSelect",function(){e.rows(a).select(),b.any()&&b.each(function(a){e.cells(a.row,a.column).select()})})}),e.on("draw.dtSelect.dt select.dtSelect.dt deselect.dtSelect.dt",function(){i(e)}),e.on("destroy.dtSelect",function(){f(e),e.off(".dtSelect")})}function k(b,c,d,e){var f=b[c+"s"]({search:"applied"}).indexes(),g=a.inArray(e,f),h=a.inArray(d,f);if(b[c+"s"]({selected:!0}).any()||-1!==g){if(g>h){var i=h;h=g,g=i}f.splice(h+1,f.length),f.splice(0,g)}else f.splice(a.inArray(d,f)+1,f.length);b[c](d,{selected:!0}).any()?(f.splice(a.inArray(d,f),1),b[c+"s"](f).deselect()):b[c+"s"](f).select()}function l(a,b){if(b||"single"===a._select.style){var c=new d.Api(a);c.rows({selected:!0}).deselect(),c.columns({selected:!0}).deselect(),c.cells({selected:!0}).deselect()}}function m(a,b,c,d,f){var g=b.select.style(),h=b[d](f,{selected:!0}).any();if("os"===g)if(a.ctrlKey||a.metaKey)b[d](f).select(!h);else if(a.shiftKey)"cell"===d?e(b,f,c._select_lastCell||null):k(b,d,f,c._select_lastCell?c._select_lastCell[d]:null);else{var i=b[d+"s"]({selected:!0});h&&1===i.flatten().length?b[d](f).deselect():(i.deselect(),b[d](f).select())}else b[d](f).select(!h)}function n(a,b){return function(c){return c.i18n("buttons."+a,b)}}d.select={},d.select.version="1.0.1",a.each([{type:"row",prop:"aoData"},{type:"column",prop:"aoColumns"}],function(a,b){d.ext.selector[b.type].push(function(a,d,e){var f,g=d.selected,h=[];if(g===c)return e;for(var i=0,j=e.length;j>i;i++)f=a[b.prop][e[i]],(g===!0&&f._select_selected===!0||g===!1&&!f._select_selected)&&h.push(e[i]);return h})}),d.ext.selector.cell.push(function(a,b,d){var e,f=b.selected,g=[];if(f===c)return d;for(var h=0,i=d.length;i>h;h++)e=a.aoData[d[h].row],(f===!0&&e._selected_cells&&e._selected_cells[d[h].column]===!0||f===!1&&(!e._selected_cells||!e._selected_cells[d[h].column]))&&g.push(d[h]);return g});var o=d.Api.register,p=d.Api.registerPlural;o("select()",function(){}),o("select.blurable()",function(a){return a===c?this.context[0]._select.blurable:this.iterator("table",function(b){b._select.blurable=a})}),o("select.info()",function(a){return i===c?this.context[0]._select.info:this.iterator("table",function(b){b._select.info=a})}),o("select.items()",function(a){return a===c?this.context[0]._select.items:this.iterator("table",function(b){b._select.items=a,h(new d.Api(b),"selectItems",[a])})}),o("select.style()",function(a){return a===c?this.context[0]._select.style:this.iterator("table",function(b){b._select.style=a,b._select_init||j(b);var c=new d.Api(b);f(c),"api"!==a&&g(c),h(new d.Api(b),"selectStyle",[a])})}),o("select.selector()",function(a){return a===c?this.context[0]._select.selector:this.iterator("table",function(b){f(new d.Api(b)),b._select.selector=a,"api"!==b._select.style&&g(new d.Api(b))})}),p("rows().select()","row().select()",function(b){var c=this;return b===!1?this.deselect():(this.iterator("row",function(b,c){l(b),b.aoData[c]._select_selected=!0,a(b.aoData[c].nTr).addClass("selected")}),this.iterator("table",function(a,b){h(c,"select",["row",c[b]],!0)}),this)}),p("columns().select()","column().select()",function(b){var c=this;return b===!1?this.deselect():(this.iterator("column",function(b,c){l(b),b.aoColumns[c]._select_selected=!0;var e=new d.Api(b).column(c);a(e.header()).addClass("selected"),a(e.footer()).addClass("selected"),e.nodes().to$().addClass("selected")}),this.iterator("table",function(a,b){h(c,"select",["column",c[b]],!0)}),this)}),p("cells().select()","cell().select()",function(b){var d=this;return b===!1?this.deselect():(this.iterator("cell",function(b,d,e){l(b);var f=b.aoData[d];f._selected_cells===c&&(f._selected_cells=[]),f._selected_cells[e]=!0,f.anCells&&a(f.anCells[e]).addClass("selected")}),this.iterator("table",function(a,b){h(d,"select",["cell",d[b]],!0)}),this)}),p("rows().deselect()","row().deselect()",function(){var b=this;return this.iterator("row",function(b,c){b.aoData[c]._select_selected=!1,a(b.aoData[c].nTr).removeClass("selected")}),this.iterator("table",function(a,c){h(b,"deselect",["row",b[c]],!0)}),this}),p("columns().deselect()","column().deselect()",function(){var b=this;return this.iterator("column",function(b,c){b.aoColumns[c]._select_selected=!1;var e=new d.Api(b),f=e.column(c);a(f.header()).removeClass("selected"),a(f.footer()).removeClass("selected"),e.cells(null,c).indexes().each(function(c){var d=b.aoData[c.row],e=d._selected_cells;!d.anCells||e&&e[c.column]||a(d.anCells[c.column]).removeClass("selected")})}),this.iterator("table",function(a,c){h(b,"deselect",["column",b[c]],!0)}),this}),p("cells().deselect()","cell().deselect()",function(){var b=this;return this.iterator("cell",function(b,c,d){var e=b.aoData[c];e._selected_cells[d]=!1,e.anCells&&!b.aoColumns[d]._select_selected&&a(e.anCells[d]).removeClass("selected")}),this.iterator("table",function(a,c){h(b,"deselect",["cell",b[c]],!0)}),this}),a.extend(d.ext.buttons,{selected:{text:n("selected","Selected"),className:"buttons-selected",init:function(a,b,c){var d=this;a.on("draw.dt.DT select.dt.DT deselect.dt.DT",function(){var a=d.rows({selected:!0}).any()||d.columns({selected:!0}).any()||d.cells({selected:!0}).any();d.enable(a)}),this.disable()}},selectedSingle:{text:n("selectedSingle","Selected single"),className:"buttons-selected-single",init:function(a,b,c){var d=this;a.on("draw.dt.DT select.dt.DT deselect.dt.DT",function(){var b=a.rows({selected:!0}).flatten().length+a.columns({selected:!0}).flatten().length+a.cells({selected:!0}).flatten().length;d.enable(1===b)}),this.disable()}},selectAll:{text:n("selectAll","Select all"),className:"buttons-select-all",action:function(){var a=this.select.items();this[a+"s"]().select()}},selectNone:{text:n("selectNone","Deselect all"),className:"buttons-select-none",action:function(){l(this.settings()[0],!0)}}}),a.each(["Row","Column","Cell"],function(a,b){var c=b.toLowerCase();d.ext.buttons["select"+b+"s"]={text:n("select"+b+"s","Select "+c+"s"),className:"buttons-select-"+c+"s",action:function(){this.select.items(c)},init:function(a,b,d){var e=this;a.on("selectItems.dt.DT",function(a,b,d){e.active(d===c)})}}}),a(b).on("init.dt.dtSelect",function(b,e,f){if("dt"===b.namespace){var g=e.oInit.select||d.defaults.select,h=new d.Api(e),i="row",j="api",k=!1,l=!0,m="td, th";e._select={},g===!0?j="os":"string"==typeof g?j=g:a.isPlainObject(g)&&(g.blurable!==c&&(k=g.blurable),g.info!==c&&(l=g.info),g.items!==c&&(i=g.items),g.style!==c&&(j=g.style),g.selector!==c&&(m=g.selector)),h.select.selector(m),h.select.items(i),h.select.style(j),h.select.blurable(k),h.select.info(l),a(h.table().node()).hasClass("selectable")&&h.select.style("os")}})};"function"==typeof define&&define.amd?define(["jquery","datatables"],d):"object"==typeof exports?d(require("jquery"),require("datatables")):jQuery&&!jQuery.fn.dataTable.select&&d(jQuery,jQuery.fn.dataTable)}(window,document);
//! moment.js
//! version : 2.10.6
//! authors : Tim Wood, Iskren Chernev, Moment.js contributors
//! license : MIT
//! momentjs.com
!function(a,b){"object"==typeof exports&&"undefined"!=typeof module?module.exports=b():"function"==typeof define&&define.amd?define(b):a.moment=b()}(this,function(){"use strict";function a(){return Hc.apply(null,arguments)}function b(a){Hc=a}function c(a){return"[object Array]"===Object.prototype.toString.call(a)}function d(a){return a instanceof Date||"[object Date]"===Object.prototype.toString.call(a)}function e(a,b){var c,d=[];for(c=0;c<a.length;++c)d.push(b(a[c],c));return d}function f(a,b){return Object.prototype.hasOwnProperty.call(a,b)}function g(a,b){for(var c in b)f(b,c)&&(a[c]=b[c]);return f(b,"toString")&&(a.toString=b.toString),f(b,"valueOf")&&(a.valueOf=b.valueOf),a}function h(a,b,c,d){return Ca(a,b,c,d,!0).utc()}function i(){return{empty:!1,unusedTokens:[],unusedInput:[],overflow:-2,charsLeftOver:0,nullInput:!1,invalidMonth:null,invalidFormat:!1,userInvalidated:!1,iso:!1}}function j(a){return null==a._pf&&(a._pf=i()),a._pf}function k(a){if(null==a._isValid){var b=j(a);a._isValid=!(isNaN(a._d.getTime())||!(b.overflow<0)||b.empty||b.invalidMonth||b.invalidWeekday||b.nullInput||b.invalidFormat||b.userInvalidated),a._strict&&(a._isValid=a._isValid&&0===b.charsLeftOver&&0===b.unusedTokens.length&&void 0===b.bigHour)}return a._isValid}function l(a){var b=h(NaN);return null!=a?g(j(b),a):j(b).userInvalidated=!0,b}function m(a,b){var c,d,e;if("undefined"!=typeof b._isAMomentObject&&(a._isAMomentObject=b._isAMomentObject),"undefined"!=typeof b._i&&(a._i=b._i),"undefined"!=typeof b._f&&(a._f=b._f),"undefined"!=typeof b._l&&(a._l=b._l),"undefined"!=typeof b._strict&&(a._strict=b._strict),"undefined"!=typeof b._tzm&&(a._tzm=b._tzm),"undefined"!=typeof b._isUTC&&(a._isUTC=b._isUTC),"undefined"!=typeof b._offset&&(a._offset=b._offset),"undefined"!=typeof b._pf&&(a._pf=j(b)),"undefined"!=typeof b._locale&&(a._locale=b._locale),Jc.length>0)for(c in Jc)d=Jc[c],e=b[d],"undefined"!=typeof e&&(a[d]=e);return a}function n(b){m(this,b),this._d=new Date(null!=b._d?b._d.getTime():NaN),Kc===!1&&(Kc=!0,a.updateOffset(this),Kc=!1)}function o(a){return a instanceof n||null!=a&&null!=a._isAMomentObject}function p(a){return 0>a?Math.ceil(a):Math.floor(a)}function q(a){var b=+a,c=0;return 0!==b&&isFinite(b)&&(c=p(b)),c}function r(a,b,c){var d,e=Math.min(a.length,b.length),f=Math.abs(a.length-b.length),g=0;for(d=0;e>d;d++)(c&&a[d]!==b[d]||!c&&q(a[d])!==q(b[d]))&&g++;return g+f}function s(){}function t(a){return a?a.toLowerCase().replace("_","-"):a}function u(a){for(var b,c,d,e,f=0;f<a.length;){for(e=t(a[f]).split("-"),b=e.length,c=t(a[f+1]),c=c?c.split("-"):null;b>0;){if(d=v(e.slice(0,b).join("-")))return d;if(c&&c.length>=b&&r(e,c,!0)>=b-1)break;b--}f++}return null}function v(a){var b=null;if(!Lc[a]&&"undefined"!=typeof module&&module&&module.exports)try{b=Ic._abbr,require("./locale/"+a),w(b)}catch(c){}return Lc[a]}function w(a,b){var c;return a&&(c="undefined"==typeof b?y(a):x(a,b),c&&(Ic=c)),Ic._abbr}function x(a,b){return null!==b?(b.abbr=a,Lc[a]=Lc[a]||new s,Lc[a].set(b),w(a),Lc[a]):(delete Lc[a],null)}function y(a){var b;if(a&&a._locale&&a._locale._abbr&&(a=a._locale._abbr),!a)return Ic;if(!c(a)){if(b=v(a))return b;a=[a]}return u(a)}function z(a,b){var c=a.toLowerCase();Mc[c]=Mc[c+"s"]=Mc[b]=a}function A(a){return"string"==typeof a?Mc[a]||Mc[a.toLowerCase()]:void 0}function B(a){var b,c,d={};for(c in a)f(a,c)&&(b=A(c),b&&(d[b]=a[c]));return d}function C(b,c){return function(d){return null!=d?(E(this,b,d),a.updateOffset(this,c),this):D(this,b)}}function D(a,b){return a._d["get"+(a._isUTC?"UTC":"")+b]()}function E(a,b,c){return a._d["set"+(a._isUTC?"UTC":"")+b](c)}function F(a,b){var c;if("object"==typeof a)for(c in a)this.set(c,a[c]);else if(a=A(a),"function"==typeof this[a])return this[a](b);return this}function G(a,b,c){var d=""+Math.abs(a),e=b-d.length,f=a>=0;return(f?c?"+":"":"-")+Math.pow(10,Math.max(0,e)).toString().substr(1)+d}function H(a,b,c,d){var e=d;"string"==typeof d&&(e=function(){return this[d]()}),a&&(Qc[a]=e),b&&(Qc[b[0]]=function(){return G(e.apply(this,arguments),b[1],b[2])}),c&&(Qc[c]=function(){return this.localeData().ordinal(e.apply(this,arguments),a)})}function I(a){return a.match(/\[[\s\S]/)?a.replace(/^\[|\]$/g,""):a.replace(/\\/g,"")}function J(a){var b,c,d=a.match(Nc);for(b=0,c=d.length;c>b;b++)Qc[d[b]]?d[b]=Qc[d[b]]:d[b]=I(d[b]);return function(e){var f="";for(b=0;c>b;b++)f+=d[b]instanceof Function?d[b].call(e,a):d[b];return f}}function K(a,b){return a.isValid()?(b=L(b,a.localeData()),Pc[b]=Pc[b]||J(b),Pc[b](a)):a.localeData().invalidDate()}function L(a,b){function c(a){return b.longDateFormat(a)||a}var d=5;for(Oc.lastIndex=0;d>=0&&Oc.test(a);)a=a.replace(Oc,c),Oc.lastIndex=0,d-=1;return a}function M(a){return"function"==typeof a&&"[object Function]"===Object.prototype.toString.call(a)}function N(a,b,c){dd[a]=M(b)?b:function(a){return a&&c?c:b}}function O(a,b){return f(dd,a)?dd[a](b._strict,b._locale):new RegExp(P(a))}function P(a){return a.replace("\\","").replace(/\\(\[)|\\(\])|\[([^\]\[]*)\]|\\(.)/g,function(a,b,c,d,e){return b||c||d||e}).replace(/[-\/\\^$*+?.()|[\]{}]/g,"\\$&")}function Q(a,b){var c,d=b;for("string"==typeof a&&(a=[a]),"number"==typeof b&&(d=function(a,c){c[b]=q(a)}),c=0;c<a.length;c++)ed[a[c]]=d}function R(a,b){Q(a,function(a,c,d,e){d._w=d._w||{},b(a,d._w,d,e)})}function S(a,b,c){null!=b&&f(ed,a)&&ed[a](b,c._a,c,a)}function T(a,b){return new Date(Date.UTC(a,b+1,0)).getUTCDate()}function U(a){return this._months[a.month()]}function V(a){return this._monthsShort[a.month()]}function W(a,b,c){var d,e,f;for(this._monthsParse||(this._monthsParse=[],this._longMonthsParse=[],this._shortMonthsParse=[]),d=0;12>d;d++){if(e=h([2e3,d]),c&&!this._longMonthsParse[d]&&(this._longMonthsParse[d]=new RegExp("^"+this.months(e,"").replace(".","")+"$","i"),this._shortMonthsParse[d]=new RegExp("^"+this.monthsShort(e,"").replace(".","")+"$","i")),c||this._monthsParse[d]||(f="^"+this.months(e,"")+"|^"+this.monthsShort(e,""),this._monthsParse[d]=new RegExp(f.replace(".",""),"i")),c&&"MMMM"===b&&this._longMonthsParse[d].test(a))return d;if(c&&"MMM"===b&&this._shortMonthsParse[d].test(a))return d;if(!c&&this._monthsParse[d].test(a))return d}}function X(a,b){var c;return"string"==typeof b&&(b=a.localeData().monthsParse(b),"number"!=typeof b)?a:(c=Math.min(a.date(),T(a.year(),b)),a._d["set"+(a._isUTC?"UTC":"")+"Month"](b,c),a)}function Y(b){return null!=b?(X(this,b),a.updateOffset(this,!0),this):D(this,"Month")}function Z(){return T(this.year(),this.month())}function $(a){var b,c=a._a;return c&&-2===j(a).overflow&&(b=c[gd]<0||c[gd]>11?gd:c[hd]<1||c[hd]>T(c[fd],c[gd])?hd:c[id]<0||c[id]>24||24===c[id]&&(0!==c[jd]||0!==c[kd]||0!==c[ld])?id:c[jd]<0||c[jd]>59?jd:c[kd]<0||c[kd]>59?kd:c[ld]<0||c[ld]>999?ld:-1,j(a)._overflowDayOfYear&&(fd>b||b>hd)&&(b=hd),j(a).overflow=b),a}function _(b){a.suppressDeprecationWarnings===!1&&"undefined"!=typeof console&&console.warn&&console.warn("Deprecation warning: "+b)}function aa(a,b){var c=!0;return g(function(){return c&&(_(a+"\n"+(new Error).stack),c=!1),b.apply(this,arguments)},b)}function ba(a,b){od[a]||(_(b),od[a]=!0)}function ca(a){var b,c,d=a._i,e=pd.exec(d);if(e){for(j(a).iso=!0,b=0,c=qd.length;c>b;b++)if(qd[b][1].exec(d)){a._f=qd[b][0];break}for(b=0,c=rd.length;c>b;b++)if(rd[b][1].exec(d)){a._f+=(e[6]||" ")+rd[b][0];break}d.match(ad)&&(a._f+="Z"),va(a)}else a._isValid=!1}function da(b){var c=sd.exec(b._i);return null!==c?void(b._d=new Date(+c[1])):(ca(b),void(b._isValid===!1&&(delete b._isValid,a.createFromInputFallback(b))))}function ea(a,b,c,d,e,f,g){var h=new Date(a,b,c,d,e,f,g);return 1970>a&&h.setFullYear(a),h}function fa(a){var b=new Date(Date.UTC.apply(null,arguments));return 1970>a&&b.setUTCFullYear(a),b}function ga(a){return ha(a)?366:365}function ha(a){return a%4===0&&a%100!==0||a%400===0}function ia(){return ha(this.year())}function ja(a,b,c){var d,e=c-b,f=c-a.day();return f>e&&(f-=7),e-7>f&&(f+=7),d=Da(a).add(f,"d"),{week:Math.ceil(d.dayOfYear()/7),year:d.year()}}function ka(a){return ja(a,this._week.dow,this._week.doy).week}function la(){return this._week.dow}function ma(){return this._week.doy}function na(a){var b=this.localeData().week(this);return null==a?b:this.add(7*(a-b),"d")}function oa(a){var b=ja(this,1,4).week;return null==a?b:this.add(7*(a-b),"d")}function pa(a,b,c,d,e){var f,g=6+e-d,h=fa(a,0,1+g),i=h.getUTCDay();return e>i&&(i+=7),c=null!=c?1*c:e,f=1+g+7*(b-1)-i+c,{year:f>0?a:a-1,dayOfYear:f>0?f:ga(a-1)+f}}function qa(a){var b=Math.round((this.clone().startOf("day")-this.clone().startOf("year"))/864e5)+1;return null==a?b:this.add(a-b,"d")}function ra(a,b,c){return null!=a?a:null!=b?b:c}function sa(a){var b=new Date;return a._useUTC?[b.getUTCFullYear(),b.getUTCMonth(),b.getUTCDate()]:[b.getFullYear(),b.getMonth(),b.getDate()]}function ta(a){var b,c,d,e,f=[];if(!a._d){for(d=sa(a),a._w&&null==a._a[hd]&&null==a._a[gd]&&ua(a),a._dayOfYear&&(e=ra(a._a[fd],d[fd]),a._dayOfYear>ga(e)&&(j(a)._overflowDayOfYear=!0),c=fa(e,0,a._dayOfYear),a._a[gd]=c.getUTCMonth(),a._a[hd]=c.getUTCDate()),b=0;3>b&&null==a._a[b];++b)a._a[b]=f[b]=d[b];for(;7>b;b++)a._a[b]=f[b]=null==a._a[b]?2===b?1:0:a._a[b];24===a._a[id]&&0===a._a[jd]&&0===a._a[kd]&&0===a._a[ld]&&(a._nextDay=!0,a._a[id]=0),a._d=(a._useUTC?fa:ea).apply(null,f),null!=a._tzm&&a._d.setUTCMinutes(a._d.getUTCMinutes()-a._tzm),a._nextDay&&(a._a[id]=24)}}function ua(a){var b,c,d,e,f,g,h;b=a._w,null!=b.GG||null!=b.W||null!=b.E?(f=1,g=4,c=ra(b.GG,a._a[fd],ja(Da(),1,4).year),d=ra(b.W,1),e=ra(b.E,1)):(f=a._locale._week.dow,g=a._locale._week.doy,c=ra(b.gg,a._a[fd],ja(Da(),f,g).year),d=ra(b.w,1),null!=b.d?(e=b.d,f>e&&++d):e=null!=b.e?b.e+f:f),h=pa(c,d,e,g,f),a._a[fd]=h.year,a._dayOfYear=h.dayOfYear}function va(b){if(b._f===a.ISO_8601)return void ca(b);b._a=[],j(b).empty=!0;var c,d,e,f,g,h=""+b._i,i=h.length,k=0;for(e=L(b._f,b._locale).match(Nc)||[],c=0;c<e.length;c++)f=e[c],d=(h.match(O(f,b))||[])[0],d&&(g=h.substr(0,h.indexOf(d)),g.length>0&&j(b).unusedInput.push(g),h=h.slice(h.indexOf(d)+d.length),k+=d.length),Qc[f]?(d?j(b).empty=!1:j(b).unusedTokens.push(f),S(f,d,b)):b._strict&&!d&&j(b).unusedTokens.push(f);j(b).charsLeftOver=i-k,h.length>0&&j(b).unusedInput.push(h),j(b).bigHour===!0&&b._a[id]<=12&&b._a[id]>0&&(j(b).bigHour=void 0),b._a[id]=wa(b._locale,b._a[id],b._meridiem),ta(b),$(b)}function wa(a,b,c){var d;return null==c?b:null!=a.meridiemHour?a.meridiemHour(b,c):null!=a.isPM?(d=a.isPM(c),d&&12>b&&(b+=12),d||12!==b||(b=0),b):b}function xa(a){var b,c,d,e,f;if(0===a._f.length)return j(a).invalidFormat=!0,void(a._d=new Date(NaN));for(e=0;e<a._f.length;e++)f=0,b=m({},a),null!=a._useUTC&&(b._useUTC=a._useUTC),b._f=a._f[e],va(b),k(b)&&(f+=j(b).charsLeftOver,f+=10*j(b).unusedTokens.length,j(b).score=f,(null==d||d>f)&&(d=f,c=b));g(a,c||b)}function ya(a){if(!a._d){var b=B(a._i);a._a=[b.year,b.month,b.day||b.date,b.hour,b.minute,b.second,b.millisecond],ta(a)}}function za(a){var b=new n($(Aa(a)));return b._nextDay&&(b.add(1,"d"),b._nextDay=void 0),b}function Aa(a){var b=a._i,e=a._f;return a._locale=a._locale||y(a._l),null===b||void 0===e&&""===b?l({nullInput:!0}):("string"==typeof b&&(a._i=b=a._locale.preparse(b)),o(b)?new n($(b)):(c(e)?xa(a):e?va(a):d(b)?a._d=b:Ba(a),a))}function Ba(b){var f=b._i;void 0===f?b._d=new Date:d(f)?b._d=new Date(+f):"string"==typeof f?da(b):c(f)?(b._a=e(f.slice(0),function(a){return parseInt(a,10)}),ta(b)):"object"==typeof f?ya(b):"number"==typeof f?b._d=new Date(f):a.createFromInputFallback(b)}function Ca(a,b,c,d,e){var f={};return"boolean"==typeof c&&(d=c,c=void 0),f._isAMomentObject=!0,f._useUTC=f._isUTC=e,f._l=c,f._i=a,f._f=b,f._strict=d,za(f)}function Da(a,b,c,d){return Ca(a,b,c,d,!1)}function Ea(a,b){var d,e;if(1===b.length&&c(b[0])&&(b=b[0]),!b.length)return Da();for(d=b[0],e=1;e<b.length;++e)(!b[e].isValid()||b[e][a](d))&&(d=b[e]);return d}function Fa(){var a=[].slice.call(arguments,0);return Ea("isBefore",a)}function Ga(){var a=[].slice.call(arguments,0);return Ea("isAfter",a)}function Ha(a){var b=B(a),c=b.year||0,d=b.quarter||0,e=b.month||0,f=b.week||0,g=b.day||0,h=b.hour||0,i=b.minute||0,j=b.second||0,k=b.millisecond||0;this._milliseconds=+k+1e3*j+6e4*i+36e5*h,this._days=+g+7*f,this._months=+e+3*d+12*c,this._data={},this._locale=y(),this._bubble()}function Ia(a){return a instanceof Ha}function Ja(a,b){H(a,0,0,function(){var a=this.utcOffset(),c="+";return 0>a&&(a=-a,c="-"),c+G(~~(a/60),2)+b+G(~~a%60,2)})}function Ka(a){var b=(a||"").match(ad)||[],c=b[b.length-1]||[],d=(c+"").match(xd)||["-",0,0],e=+(60*d[1])+q(d[2]);return"+"===d[0]?e:-e}function La(b,c){var e,f;return c._isUTC?(e=c.clone(),f=(o(b)||d(b)?+b:+Da(b))-+e,e._d.setTime(+e._d+f),a.updateOffset(e,!1),e):Da(b).local()}function Ma(a){return 15*-Math.round(a._d.getTimezoneOffset()/15)}function Na(b,c){var d,e=this._offset||0;return null!=b?("string"==typeof b&&(b=Ka(b)),Math.abs(b)<16&&(b=60*b),!this._isUTC&&c&&(d=Ma(this)),this._offset=b,this._isUTC=!0,null!=d&&this.add(d,"m"),e!==b&&(!c||this._changeInProgress?bb(this,Ya(b-e,"m"),1,!1):this._changeInProgress||(this._changeInProgress=!0,a.updateOffset(this,!0),this._changeInProgress=null)),this):this._isUTC?e:Ma(this)}function Oa(a,b){return null!=a?("string"!=typeof a&&(a=-a),this.utcOffset(a,b),this):-this.utcOffset()}function Pa(a){return this.utcOffset(0,a)}function Qa(a){return this._isUTC&&(this.utcOffset(0,a),this._isUTC=!1,a&&this.subtract(Ma(this),"m")),this}function Ra(){return this._tzm?this.utcOffset(this._tzm):"string"==typeof this._i&&this.utcOffset(Ka(this._i)),this}function Sa(a){return a=a?Da(a).utcOffset():0,(this.utcOffset()-a)%60===0}function Ta(){return this.utcOffset()>this.clone().month(0).utcOffset()||this.utcOffset()>this.clone().month(5).utcOffset()}function Ua(){if("undefined"!=typeof this._isDSTShifted)return this._isDSTShifted;var a={};if(m(a,this),a=Aa(a),a._a){var b=a._isUTC?h(a._a):Da(a._a);this._isDSTShifted=this.isValid()&&r(a._a,b.toArray())>0}else this._isDSTShifted=!1;return this._isDSTShifted}function Va(){return!this._isUTC}function Wa(){return this._isUTC}function Xa(){return this._isUTC&&0===this._offset}function Ya(a,b){var c,d,e,g=a,h=null;return Ia(a)?g={ms:a._milliseconds,d:a._days,M:a._months}:"number"==typeof a?(g={},b?g[b]=a:g.milliseconds=a):(h=yd.exec(a))?(c="-"===h[1]?-1:1,g={y:0,d:q(h[hd])*c,h:q(h[id])*c,m:q(h[jd])*c,s:q(h[kd])*c,ms:q(h[ld])*c}):(h=zd.exec(a))?(c="-"===h[1]?-1:1,g={y:Za(h[2],c),M:Za(h[3],c),d:Za(h[4],c),h:Za(h[5],c),m:Za(h[6],c),s:Za(h[7],c),w:Za(h[8],c)}):null==g?g={}:"object"==typeof g&&("from"in g||"to"in g)&&(e=_a(Da(g.from),Da(g.to)),g={},g.ms=e.milliseconds,g.M=e.months),d=new Ha(g),Ia(a)&&f(a,"_locale")&&(d._locale=a._locale),d}function Za(a,b){var c=a&&parseFloat(a.replace(",","."));return(isNaN(c)?0:c)*b}function $a(a,b){var c={milliseconds:0,months:0};return c.months=b.month()-a.month()+12*(b.year()-a.year()),a.clone().add(c.months,"M").isAfter(b)&&--c.months,c.milliseconds=+b-+a.clone().add(c.months,"M"),c}function _a(a,b){var c;return b=La(b,a),a.isBefore(b)?c=$a(a,b):(c=$a(b,a),c.milliseconds=-c.milliseconds,c.months=-c.months),c}function ab(a,b){return function(c,d){var e,f;return null===d||isNaN(+d)||(ba(b,"moment()."+b+"(period, number) is deprecated. Please use moment()."+b+"(number, period)."),f=c,c=d,d=f),c="string"==typeof c?+c:c,e=Ya(c,d),bb(this,e,a),this}}function bb(b,c,d,e){var f=c._milliseconds,g=c._days,h=c._months;e=null==e?!0:e,f&&b._d.setTime(+b._d+f*d),g&&E(b,"Date",D(b,"Date")+g*d),h&&X(b,D(b,"Month")+h*d),e&&a.updateOffset(b,g||h)}function cb(a,b){var c=a||Da(),d=La(c,this).startOf("day"),e=this.diff(d,"days",!0),f=-6>e?"sameElse":-1>e?"lastWeek":0>e?"lastDay":1>e?"sameDay":2>e?"nextDay":7>e?"nextWeek":"sameElse";return this.format(b&&b[f]||this.localeData().calendar(f,this,Da(c)))}function db(){return new n(this)}function eb(a,b){var c;return b=A("undefined"!=typeof b?b:"millisecond"),"millisecond"===b?(a=o(a)?a:Da(a),+this>+a):(c=o(a)?+a:+Da(a),c<+this.clone().startOf(b))}function fb(a,b){var c;return b=A("undefined"!=typeof b?b:"millisecond"),"millisecond"===b?(a=o(a)?a:Da(a),+a>+this):(c=o(a)?+a:+Da(a),+this.clone().endOf(b)<c)}function gb(a,b,c){return this.isAfter(a,c)&&this.isBefore(b,c)}function hb(a,b){var c;return b=A(b||"millisecond"),"millisecond"===b?(a=o(a)?a:Da(a),+this===+a):(c=+Da(a),+this.clone().startOf(b)<=c&&c<=+this.clone().endOf(b))}function ib(a,b,c){var d,e,f=La(a,this),g=6e4*(f.utcOffset()-this.utcOffset());return b=A(b),"year"===b||"month"===b||"quarter"===b?(e=jb(this,f),"quarter"===b?e/=3:"year"===b&&(e/=12)):(d=this-f,e="second"===b?d/1e3:"minute"===b?d/6e4:"hour"===b?d/36e5:"day"===b?(d-g)/864e5:"week"===b?(d-g)/6048e5:d),c?e:p(e)}function jb(a,b){var c,d,e=12*(b.year()-a.year())+(b.month()-a.month()),f=a.clone().add(e,"months");return 0>b-f?(c=a.clone().add(e-1,"months"),d=(b-f)/(f-c)):(c=a.clone().add(e+1,"months"),d=(b-f)/(c-f)),-(e+d)}function kb(){return this.clone().locale("en").format("ddd MMM DD YYYY HH:mm:ss [GMT]ZZ")}function lb(){var a=this.clone().utc();return 0<a.year()&&a.year()<=9999?"function"==typeof Date.prototype.toISOString?this.toDate().toISOString():K(a,"YYYY-MM-DD[T]HH:mm:ss.SSS[Z]"):K(a,"YYYYYY-MM-DD[T]HH:mm:ss.SSS[Z]")}function mb(b){var c=K(this,b||a.defaultFormat);return this.localeData().postformat(c)}function nb(a,b){return this.isValid()?Ya({to:this,from:a}).locale(this.locale()).humanize(!b):this.localeData().invalidDate()}function ob(a){return this.from(Da(),a)}function pb(a,b){return this.isValid()?Ya({from:this,to:a}).locale(this.locale()).humanize(!b):this.localeData().invalidDate()}function qb(a){return this.to(Da(),a)}function rb(a){var b;return void 0===a?this._locale._abbr:(b=y(a),null!=b&&(this._locale=b),this)}function sb(){return this._locale}function tb(a){switch(a=A(a)){case"year":this.month(0);case"quarter":case"month":this.date(1);case"week":case"isoWeek":case"day":this.hours(0);case"hour":this.minutes(0);case"minute":this.seconds(0);case"second":this.milliseconds(0)}return"week"===a&&this.weekday(0),"isoWeek"===a&&this.isoWeekday(1),"quarter"===a&&this.month(3*Math.floor(this.month()/3)),this}function ub(a){return a=A(a),void 0===a||"millisecond"===a?this:this.startOf(a).add(1,"isoWeek"===a?"week":a).subtract(1,"ms")}function vb(){return+this._d-6e4*(this._offset||0)}function wb(){return Math.floor(+this/1e3)}function xb(){return this._offset?new Date(+this):this._d}function yb(){var a=this;return[a.year(),a.month(),a.date(),a.hour(),a.minute(),a.second(),a.millisecond()]}function zb(){var a=this;return{years:a.year(),months:a.month(),date:a.date(),hours:a.hours(),minutes:a.minutes(),seconds:a.seconds(),milliseconds:a.milliseconds()}}function Ab(){return k(this)}function Bb(){return g({},j(this))}function Cb(){return j(this).overflow}function Db(a,b){H(0,[a,a.length],0,b)}function Eb(a,b,c){return ja(Da([a,11,31+b-c]),b,c).week}function Fb(a){var b=ja(this,this.localeData()._week.dow,this.localeData()._week.doy).year;return null==a?b:this.add(a-b,"y")}function Gb(a){var b=ja(this,1,4).year;return null==a?b:this.add(a-b,"y")}function Hb(){return Eb(this.year(),1,4)}function Ib(){var a=this.localeData()._week;return Eb(this.year(),a.dow,a.doy)}function Jb(a){return null==a?Math.ceil((this.month()+1)/3):this.month(3*(a-1)+this.month()%3)}function Kb(a,b){return"string"!=typeof a?a:isNaN(a)?(a=b.weekdaysParse(a),"number"==typeof a?a:null):parseInt(a,10)}function Lb(a){return this._weekdays[a.day()]}function Mb(a){return this._weekdaysShort[a.day()]}function Nb(a){return this._weekdaysMin[a.day()]}function Ob(a){var b,c,d;for(this._weekdaysParse=this._weekdaysParse||[],b=0;7>b;b++)if(this._weekdaysParse[b]||(c=Da([2e3,1]).day(b),d="^"+this.weekdays(c,"")+"|^"+this.weekdaysShort(c,"")+"|^"+this.weekdaysMin(c,""),this._weekdaysParse[b]=new RegExp(d.replace(".",""),"i")),this._weekdaysParse[b].test(a))return b}function Pb(a){var b=this._isUTC?this._d.getUTCDay():this._d.getDay();return null!=a?(a=Kb(a,this.localeData()),this.add(a-b,"d")):b}function Qb(a){var b=(this.day()+7-this.localeData()._week.dow)%7;return null==a?b:this.add(a-b,"d")}function Rb(a){return null==a?this.day()||7:this.day(this.day()%7?a:a-7)}function Sb(a,b){H(a,0,0,function(){return this.localeData().meridiem(this.hours(),this.minutes(),b)})}function Tb(a,b){return b._meridiemParse}function Ub(a){return"p"===(a+"").toLowerCase().charAt(0)}function Vb(a,b,c){return a>11?c?"pm":"PM":c?"am":"AM"}function Wb(a,b){b[ld]=q(1e3*("0."+a))}function Xb(){return this._isUTC?"UTC":""}function Yb(){return this._isUTC?"Coordinated Universal Time":""}function Zb(a){return Da(1e3*a)}function $b(){return Da.apply(null,arguments).parseZone()}function _b(a,b,c){var d=this._calendar[a];return"function"==typeof d?d.call(b,c):d}function ac(a){var b=this._longDateFormat[a],c=this._longDateFormat[a.toUpperCase()];return b||!c?b:(this._longDateFormat[a]=c.replace(/MMMM|MM|DD|dddd/g,function(a){return a.slice(1)}),this._longDateFormat[a])}function bc(){return this._invalidDate}function cc(a){return this._ordinal.replace("%d",a)}function dc(a){return a}function ec(a,b,c,d){var e=this._relativeTime[c];return"function"==typeof e?e(a,b,c,d):e.replace(/%d/i,a)}function fc(a,b){var c=this._relativeTime[a>0?"future":"past"];return"function"==typeof c?c(b):c.replace(/%s/i,b)}function gc(a){var b,c;for(c in a)b=a[c],"function"==typeof b?this[c]=b:this["_"+c]=b;this._ordinalParseLenient=new RegExp(this._ordinalParse.source+"|"+/\d{1,2}/.source)}function hc(a,b,c,d){var e=y(),f=h().set(d,b);return e[c](f,a)}function ic(a,b,c,d,e){if("number"==typeof a&&(b=a,a=void 0),a=a||"",null!=b)return hc(a,b,c,e);var f,g=[];for(f=0;d>f;f++)g[f]=hc(a,f,c,e);return g}function jc(a,b){return ic(a,b,"months",12,"month")}function kc(a,b){return ic(a,b,"monthsShort",12,"month")}function lc(a,b){return ic(a,b,"weekdays",7,"day")}function mc(a,b){return ic(a,b,"weekdaysShort",7,"day")}function nc(a,b){return ic(a,b,"weekdaysMin",7,"day")}function oc(){var a=this._data;return this._milliseconds=Wd(this._milliseconds),this._days=Wd(this._days),this._months=Wd(this._months),a.milliseconds=Wd(a.milliseconds),a.seconds=Wd(a.seconds),a.minutes=Wd(a.minutes),a.hours=Wd(a.hours),a.months=Wd(a.months),a.years=Wd(a.years),this}function pc(a,b,c,d){var e=Ya(b,c);return a._milliseconds+=d*e._milliseconds,a._days+=d*e._days,a._months+=d*e._months,a._bubble()}function qc(a,b){return pc(this,a,b,1)}function rc(a,b){return pc(this,a,b,-1)}function sc(a){return 0>a?Math.floor(a):Math.ceil(a)}function tc(){var a,b,c,d,e,f=this._milliseconds,g=this._days,h=this._months,i=this._data;return f>=0&&g>=0&&h>=0||0>=f&&0>=g&&0>=h||(f+=864e5*sc(vc(h)+g),g=0,h=0),i.milliseconds=f%1e3,a=p(f/1e3),i.seconds=a%60,b=p(a/60),i.minutes=b%60,c=p(b/60),i.hours=c%24,g+=p(c/24),e=p(uc(g)),h+=e,g-=sc(vc(e)),d=p(h/12),h%=12,i.days=g,i.months=h,i.years=d,this}function uc(a){return 4800*a/146097}function vc(a){return 146097*a/4800}function wc(a){var b,c,d=this._milliseconds;if(a=A(a),"month"===a||"year"===a)return b=this._days+d/864e5,c=this._months+uc(b),"month"===a?c:c/12;switch(b=this._days+Math.round(vc(this._months)),a){case"week":return b/7+d/6048e5;case"day":return b+d/864e5;case"hour":return 24*b+d/36e5;case"minute":return 1440*b+d/6e4;case"second":return 86400*b+d/1e3;case"millisecond":return Math.floor(864e5*b)+d;default:throw new Error("Unknown unit "+a)}}function xc(){return this._milliseconds+864e5*this._days+this._months%12*2592e6+31536e6*q(this._months/12)}function yc(a){return function(){return this.as(a)}}function zc(a){return a=A(a),this[a+"s"]()}function Ac(a){return function(){return this._data[a]}}function Bc(){return p(this.days()/7)}function Cc(a,b,c,d,e){return e.relativeTime(b||1,!!c,a,d)}function Dc(a,b,c){var d=Ya(a).abs(),e=ke(d.as("s")),f=ke(d.as("m")),g=ke(d.as("h")),h=ke(d.as("d")),i=ke(d.as("M")),j=ke(d.as("y")),k=e<le.s&&["s",e]||1===f&&["m"]||f<le.m&&["mm",f]||1===g&&["h"]||g<le.h&&["hh",g]||1===h&&["d"]||h<le.d&&["dd",h]||1===i&&["M"]||i<le.M&&["MM",i]||1===j&&["y"]||["yy",j];return k[2]=b,k[3]=+a>0,k[4]=c,Cc.apply(null,k)}function Ec(a,b){return void 0===le[a]?!1:void 0===b?le[a]:(le[a]=b,!0)}function Fc(a){var b=this.localeData(),c=Dc(this,!a,b);return a&&(c=b.pastFuture(+this,c)),b.postformat(c)}function Gc(){var a,b,c,d=me(this._milliseconds)/1e3,e=me(this._days),f=me(this._months);a=p(d/60),b=p(a/60),d%=60,a%=60,c=p(f/12),f%=12;var g=c,h=f,i=e,j=b,k=a,l=d,m=this.asSeconds();return m?(0>m?"-":"")+"P"+(g?g+"Y":"")+(h?h+"M":"")+(i?i+"D":"")+(j||k||l?"T":"")+(j?j+"H":"")+(k?k+"M":"")+(l?l+"S":""):"P0D"}var Hc,Ic,Jc=a.momentProperties=[],Kc=!1,Lc={},Mc={},Nc=/(\[[^\[]*\])|(\\)?(Mo|MM?M?M?|Do|DDDo|DD?D?D?|ddd?d?|do?|w[o|w]?|W[o|W]?|Q|YYYYYY|YYYYY|YYYY|YY|gg(ggg?)?|GG(GGG?)?|e|E|a|A|hh?|HH?|mm?|ss?|S{1,9}|x|X|zz?|ZZ?|.)/g,Oc=/(\[[^\[]*\])|(\\)?(LTS|LT|LL?L?L?|l{1,4})/g,Pc={},Qc={},Rc=/\d/,Sc=/\d\d/,Tc=/\d{3}/,Uc=/\d{4}/,Vc=/[+-]?\d{6}/,Wc=/\d\d?/,Xc=/\d{1,3}/,Yc=/\d{1,4}/,Zc=/[+-]?\d{1,6}/,$c=/\d+/,_c=/[+-]?\d+/,ad=/Z|[+-]\d\d:?\d\d/gi,bd=/[+-]?\d+(\.\d{1,3})?/,cd=/[0-9]*['a-z\u00A0-\u05FF\u0700-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+|[\u0600-\u06FF\/]+(\s*?[\u0600-\u06FF]+){1,2}/i,dd={},ed={},fd=0,gd=1,hd=2,id=3,jd=4,kd=5,ld=6;H("M",["MM",2],"Mo",function(){return this.month()+1}),H("MMM",0,0,function(a){return this.localeData().monthsShort(this,a)}),H("MMMM",0,0,function(a){return this.localeData().months(this,a)}),z("month","M"),N("M",Wc),N("MM",Wc,Sc),N("MMM",cd),N("MMMM",cd),Q(["M","MM"],function(a,b){b[gd]=q(a)-1}),Q(["MMM","MMMM"],function(a,b,c,d){var e=c._locale.monthsParse(a,d,c._strict);null!=e?b[gd]=e:j(c).invalidMonth=a});var md="January_February_March_April_May_June_July_August_September_October_November_December".split("_"),nd="Jan_Feb_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec".split("_"),od={};a.suppressDeprecationWarnings=!1;var pd=/^\s*(?:[+-]\d{6}|\d{4})-(?:(\d\d-\d\d)|(W\d\d$)|(W\d\d-\d)|(\d\d\d))((T| )(\d\d(:\d\d(:\d\d(\.\d+)?)?)?)?([\+\-]\d\d(?::?\d\d)?|\s*Z)?)?$/,qd=[["YYYYYY-MM-DD",/[+-]\d{6}-\d{2}-\d{2}/],["YYYY-MM-DD",/\d{4}-\d{2}-\d{2}/],["GGGG-[W]WW-E",/\d{4}-W\d{2}-\d/],["GGGG-[W]WW",/\d{4}-W\d{2}/],["YYYY-DDD",/\d{4}-\d{3}/]],rd=[["HH:mm:ss.SSSS",/(T| )\d\d:\d\d:\d\d\.\d+/],["HH:mm:ss",/(T| )\d\d:\d\d:\d\d/],["HH:mm",/(T| )\d\d:\d\d/],["HH",/(T| )\d\d/]],sd=/^\/?Date\((\-?\d+)/i;a.createFromInputFallback=aa("moment construction falls back to js Date. This is discouraged and will be removed in upcoming major release. Please refer to https://github.com/moment/moment/issues/1407 for more info.",function(a){a._d=new Date(a._i+(a._useUTC?" UTC":""))}),H(0,["YY",2],0,function(){return this.year()%100}),H(0,["YYYY",4],0,"year"),H(0,["YYYYY",5],0,"year"),H(0,["YYYYYY",6,!0],0,"year"),z("year","y"),N("Y",_c),N("YY",Wc,Sc),N("YYYY",Yc,Uc),N("YYYYY",Zc,Vc),N("YYYYYY",Zc,Vc),Q(["YYYYY","YYYYYY"],fd),Q("YYYY",function(b,c){c[fd]=2===b.length?a.parseTwoDigitYear(b):q(b)}),Q("YY",function(b,c){c[fd]=a.parseTwoDigitYear(b)}),a.parseTwoDigitYear=function(a){return q(a)+(q(a)>68?1900:2e3)};var td=C("FullYear",!1);H("w",["ww",2],"wo","week"),H("W",["WW",2],"Wo","isoWeek"),z("week","w"),z("isoWeek","W"),N("w",Wc),N("ww",Wc,Sc),N("W",Wc),N("WW",Wc,Sc),R(["w","ww","W","WW"],function(a,b,c,d){b[d.substr(0,1)]=q(a)});var ud={dow:0,doy:6};H("DDD",["DDDD",3],"DDDo","dayOfYear"),z("dayOfYear","DDD"),N("DDD",Xc),N("DDDD",Tc),Q(["DDD","DDDD"],function(a,b,c){c._dayOfYear=q(a)}),a.ISO_8601=function(){};var vd=aa("moment().min is deprecated, use moment.min instead. https://github.com/moment/moment/issues/1548",function(){var a=Da.apply(null,arguments);return this>a?this:a}),wd=aa("moment().max is deprecated, use moment.max instead. https://github.com/moment/moment/issues/1548",function(){var a=Da.apply(null,arguments);return a>this?this:a});Ja("Z",":"),Ja("ZZ",""),N("Z",ad),N("ZZ",ad),Q(["Z","ZZ"],function(a,b,c){c._useUTC=!0,c._tzm=Ka(a)});var xd=/([\+\-]|\d\d)/gi;a.updateOffset=function(){};var yd=/(\-)?(?:(\d*)\.)?(\d+)\:(\d+)(?:\:(\d+)\.?(\d{3})?)?/,zd=/^(-)?P(?:(?:([0-9,.]*)Y)?(?:([0-9,.]*)M)?(?:([0-9,.]*)D)?(?:T(?:([0-9,.]*)H)?(?:([0-9,.]*)M)?(?:([0-9,.]*)S)?)?|([0-9,.]*)W)$/;Ya.fn=Ha.prototype;var Ad=ab(1,"add"),Bd=ab(-1,"subtract");a.defaultFormat="YYYY-MM-DDTHH:mm:ssZ";var Cd=aa("moment().lang() is deprecated. Instead, use moment().localeData() to get the language configuration. Use moment().locale() to change languages.",function(a){return void 0===a?this.localeData():this.locale(a)});H(0,["gg",2],0,function(){return this.weekYear()%100}),H(0,["GG",2],0,function(){return this.isoWeekYear()%100}),Db("gggg","weekYear"),Db("ggggg","weekYear"),Db("GGGG","isoWeekYear"),Db("GGGGG","isoWeekYear"),z("weekYear","gg"),z("isoWeekYear","GG"),N("G",_c),N("g",_c),N("GG",Wc,Sc),N("gg",Wc,Sc),N("GGGG",Yc,Uc),N("gggg",Yc,Uc),N("GGGGG",Zc,Vc),N("ggggg",Zc,Vc),R(["gggg","ggggg","GGGG","GGGGG"],function(a,b,c,d){b[d.substr(0,2)]=q(a)}),R(["gg","GG"],function(b,c,d,e){c[e]=a.parseTwoDigitYear(b)}),H("Q",0,0,"quarter"),z("quarter","Q"),N("Q",Rc),Q("Q",function(a,b){b[gd]=3*(q(a)-1)}),H("D",["DD",2],"Do","date"),z("date","D"),N("D",Wc),N("DD",Wc,Sc),N("Do",function(a,b){return a?b._ordinalParse:b._ordinalParseLenient}),Q(["D","DD"],hd),Q("Do",function(a,b){b[hd]=q(a.match(Wc)[0],10)});var Dd=C("Date",!0);H("d",0,"do","day"),H("dd",0,0,function(a){return this.localeData().weekdaysMin(this,a)}),H("ddd",0,0,function(a){return this.localeData().weekdaysShort(this,a)}),H("dddd",0,0,function(a){return this.localeData().weekdays(this,a)}),H("e",0,0,"weekday"),H("E",0,0,"isoWeekday"),z("day","d"),z("weekday","e"),z("isoWeekday","E"),N("d",Wc),N("e",Wc),N("E",Wc),N("dd",cd),N("ddd",cd),N("dddd",cd),R(["dd","ddd","dddd"],function(a,b,c){var d=c._locale.weekdaysParse(a);null!=d?b.d=d:j(c).invalidWeekday=a}),R(["d","e","E"],function(a,b,c,d){b[d]=q(a)});var Ed="Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),Fd="Sun_Mon_Tue_Wed_Thu_Fri_Sat".split("_"),Gd="Su_Mo_Tu_We_Th_Fr_Sa".split("_");H("H",["HH",2],0,"hour"),H("h",["hh",2],0,function(){return this.hours()%12||12}),Sb("a",!0),Sb("A",!1),z("hour","h"),N("a",Tb),N("A",Tb),N("H",Wc),N("h",Wc),N("HH",Wc,Sc),N("hh",Wc,Sc),Q(["H","HH"],id),Q(["a","A"],function(a,b,c){c._isPm=c._locale.isPM(a),c._meridiem=a}),Q(["h","hh"],function(a,b,c){b[id]=q(a),j(c).bigHour=!0});var Hd=/[ap]\.?m?\.?/i,Id=C("Hours",!0);H("m",["mm",2],0,"minute"),z("minute","m"),N("m",Wc),N("mm",Wc,Sc),Q(["m","mm"],jd);var Jd=C("Minutes",!1);H("s",["ss",2],0,"second"),z("second","s"),N("s",Wc),N("ss",Wc,Sc),Q(["s","ss"],kd);var Kd=C("Seconds",!1);H("S",0,0,function(){return~~(this.millisecond()/100)}),H(0,["SS",2],0,function(){return~~(this.millisecond()/10)}),H(0,["SSS",3],0,"millisecond"),H(0,["SSSS",4],0,function(){return 10*this.millisecond()}),H(0,["SSSSS",5],0,function(){return 100*this.millisecond()}),H(0,["SSSSSS",6],0,function(){return 1e3*this.millisecond()}),H(0,["SSSSSSS",7],0,function(){return 1e4*this.millisecond()}),H(0,["SSSSSSSS",8],0,function(){return 1e5*this.millisecond()}),H(0,["SSSSSSSSS",9],0,function(){return 1e6*this.millisecond()}),z("millisecond","ms"),N("S",Xc,Rc),N("SS",Xc,Sc),N("SSS",Xc,Tc);var Ld;for(Ld="SSSS";Ld.length<=9;Ld+="S")N(Ld,$c);for(Ld="S";Ld.length<=9;Ld+="S")Q(Ld,Wb);var Md=C("Milliseconds",!1);H("z",0,0,"zoneAbbr"),H("zz",0,0,"zoneName");var Nd=n.prototype;Nd.add=Ad,Nd.calendar=cb,Nd.clone=db,Nd.diff=ib,Nd.endOf=ub,Nd.format=mb,Nd.from=nb,Nd.fromNow=ob,Nd.to=pb,Nd.toNow=qb,Nd.get=F,Nd.invalidAt=Cb,Nd.isAfter=eb,Nd.isBefore=fb,Nd.isBetween=gb,Nd.isSame=hb,Nd.isValid=Ab,Nd.lang=Cd,Nd.locale=rb,Nd.localeData=sb,Nd.max=wd,Nd.min=vd,Nd.parsingFlags=Bb,Nd.set=F,Nd.startOf=tb,Nd.subtract=Bd,Nd.toArray=yb,Nd.toObject=zb,Nd.toDate=xb,Nd.toISOString=lb,Nd.toJSON=lb,Nd.toString=kb,Nd.unix=wb,Nd.valueOf=vb,Nd.year=td,Nd.isLeapYear=ia,Nd.weekYear=Fb,Nd.isoWeekYear=Gb,Nd.quarter=Nd.quarters=Jb,Nd.month=Y,Nd.daysInMonth=Z,Nd.week=Nd.weeks=na,Nd.isoWeek=Nd.isoWeeks=oa,Nd.weeksInYear=Ib,Nd.isoWeeksInYear=Hb,Nd.date=Dd,Nd.day=Nd.days=Pb,Nd.weekday=Qb,Nd.isoWeekday=Rb,Nd.dayOfYear=qa,Nd.hour=Nd.hours=Id,Nd.minute=Nd.minutes=Jd,Nd.second=Nd.seconds=Kd,
Nd.millisecond=Nd.milliseconds=Md,Nd.utcOffset=Na,Nd.utc=Pa,Nd.local=Qa,Nd.parseZone=Ra,Nd.hasAlignedHourOffset=Sa,Nd.isDST=Ta,Nd.isDSTShifted=Ua,Nd.isLocal=Va,Nd.isUtcOffset=Wa,Nd.isUtc=Xa,Nd.isUTC=Xa,Nd.zoneAbbr=Xb,Nd.zoneName=Yb,Nd.dates=aa("dates accessor is deprecated. Use date instead.",Dd),Nd.months=aa("months accessor is deprecated. Use month instead",Y),Nd.years=aa("years accessor is deprecated. Use year instead",td),Nd.zone=aa("moment().zone is deprecated, use moment().utcOffset instead. https://github.com/moment/moment/issues/1779",Oa);var Od=Nd,Pd={sameDay:"[Today at] LT",nextDay:"[Tomorrow at] LT",nextWeek:"dddd [at] LT",lastDay:"[Yesterday at] LT",lastWeek:"[Last] dddd [at] LT",sameElse:"L"},Qd={LTS:"h:mm:ss A",LT:"h:mm A",L:"MM/DD/YYYY",LL:"MMMM D, YYYY",LLL:"MMMM D, YYYY h:mm A",LLLL:"dddd, MMMM D, YYYY h:mm A"},Rd="Invalid date",Sd="%d",Td=/\d{1,2}/,Ud={future:"in %s",past:"%s ago",s:"a few seconds",m:"a minute",mm:"%d minutes",h:"an hour",hh:"%d hours",d:"a day",dd:"%d days",M:"a month",MM:"%d months",y:"a year",yy:"%d years"},Vd=s.prototype;Vd._calendar=Pd,Vd.calendar=_b,Vd._longDateFormat=Qd,Vd.longDateFormat=ac,Vd._invalidDate=Rd,Vd.invalidDate=bc,Vd._ordinal=Sd,Vd.ordinal=cc,Vd._ordinalParse=Td,Vd.preparse=dc,Vd.postformat=dc,Vd._relativeTime=Ud,Vd.relativeTime=ec,Vd.pastFuture=fc,Vd.set=gc,Vd.months=U,Vd._months=md,Vd.monthsShort=V,Vd._monthsShort=nd,Vd.monthsParse=W,Vd.week=ka,Vd._week=ud,Vd.firstDayOfYear=ma,Vd.firstDayOfWeek=la,Vd.weekdays=Lb,Vd._weekdays=Ed,Vd.weekdaysMin=Nb,Vd._weekdaysMin=Gd,Vd.weekdaysShort=Mb,Vd._weekdaysShort=Fd,Vd.weekdaysParse=Ob,Vd.isPM=Ub,Vd._meridiemParse=Hd,Vd.meridiem=Vb,w("en",{ordinalParse:/\d{1,2}(th|st|nd|rd)/,ordinal:function(a){var b=a%10,c=1===q(a%100/10)?"th":1===b?"st":2===b?"nd":3===b?"rd":"th";return a+c}}),a.lang=aa("moment.lang is deprecated. Use moment.locale instead.",w),a.langData=aa("moment.langData is deprecated. Use moment.localeData instead.",y);var Wd=Math.abs,Xd=yc("ms"),Yd=yc("s"),Zd=yc("m"),$d=yc("h"),_d=yc("d"),ae=yc("w"),be=yc("M"),ce=yc("y"),de=Ac("milliseconds"),ee=Ac("seconds"),fe=Ac("minutes"),ge=Ac("hours"),he=Ac("days"),ie=Ac("months"),je=Ac("years"),ke=Math.round,le={s:45,m:45,h:22,d:26,M:11},me=Math.abs,ne=Ha.prototype;ne.abs=oc,ne.add=qc,ne.subtract=rc,ne.as=wc,ne.asMilliseconds=Xd,ne.asSeconds=Yd,ne.asMinutes=Zd,ne.asHours=$d,ne.asDays=_d,ne.asWeeks=ae,ne.asMonths=be,ne.asYears=ce,ne.valueOf=xc,ne._bubble=tc,ne.get=zc,ne.milliseconds=de,ne.seconds=ee,ne.minutes=fe,ne.hours=ge,ne.days=he,ne.weeks=Bc,ne.months=ie,ne.years=je,ne.humanize=Fc,ne.toISOString=Gc,ne.toString=Gc,ne.toJSON=Gc,ne.locale=rb,ne.localeData=sb,ne.toIsoString=aa("toIsoString() is deprecated. Please use toISOString() instead (notice the capitals)",Gc),ne.lang=Cd,H("X",0,0,"unix"),H("x",0,0,"valueOf"),N("x",_c),N("X",bd),Q("X",function(a,b,c){c._d=new Date(1e3*parseFloat(a,10))}),Q("x",function(a,b,c){c._d=new Date(q(a))}),a.version="2.10.6",b(Da),a.fn=Od,a.min=Fa,a.max=Ga,a.utc=h,a.unix=Zb,a.months=jc,a.isDate=d,a.locale=w,a.invalid=l,a.duration=Ya,a.isMoment=o,a.weekdays=lc,a.parseZone=$b,a.localeData=y,a.isDuration=Ia,a.monthsShort=kc,a.weekdaysMin=nc,a.defineLocale=x,a.weekdaysShort=mc,a.normalizeUnits=A,a.relativeTimeThreshold=Ec;var oe=a;return oe});
﻿/*! FixedHeader 2.1.2
 * 漏2010-2014 SpryMedia Ltd - datatables.net/license
 */

/**
 * @summary     FixedHeader
 * @description Fix a table's header or footer, so it is always visible while
 *              Scrolling
 * @version     2.1.2
 * @file        dataTables.fixedHeader.js
 * @author      SpryMedia Ltd (www.sprymedia.co.uk)
 * @contact     www.sprymedia.co.uk/contact
 * @copyright   Copyright 2009-2014 SpryMedia Ltd.
 *
 * This source file is free software, available under the following license:
 *   MIT license - http://datatables.net/license/mit
 *
 * This source file is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE. See the license files for details.
 *
 * For details please refer to: http://www.datatables.net
 */

/* Global scope for FixedColumns for backwards compatibility - will be removed
 * in future. Not documented in 1.1.x.
 */

/* Global scope for FixedColumns */
var FixedHeader;

(function (window, document, undefined) {


    var factory = function ($, DataTable) {
        "use strict";

        /*
         * Function: FixedHeader
         * Purpose:  Provide 'fixed' header, footer and columns for a DataTable
         * Returns:  object:FixedHeader - must be called with 'new'
         * Inputs:   mixed:mTable - target table
         *  @param {object} dt DataTables instance or HTML table node. With DataTables
         *    1.10 this can also be a jQuery collection (with just a single table in its
         *    result set), a jQuery selector, DataTables API instance or settings
         *    object.
         *  @param {object} [oInit] initialisation settings, with the following
         *    properties (each optional)
         *    * bool:top -    fix the header (default true)
         *    * bool:bottom - fix the footer (default false)
         *    * int:left -    fix the left column(s) (default 0)
         *    * int:right -   fix the right column(s) (default 0)
         *    * int:zTop -    fixed header zIndex
         *    * int:zBottom - fixed footer zIndex
         *    * int:zLeft -   fixed left zIndex
         *    * int:zRight -  fixed right zIndex
         */
        FixedHeader = function (mTable, oInit) {
            /* Sanity check - you just know it will happen */
            if (!this instanceof FixedHeader) {
                alert("FixedHeader warning: FixedHeader must be initialised with the 'new' keyword.");
                return;
            }

            var that = this;
            var oSettings = {
                "aoCache": [],
                "oSides": {
                    "top": true,
                    "bottom": false,
                    "left": 0,
                    "right": 0
                },
                "oZIndexes": {
                    "top": 104,
                    "bottom": 103,
                    "left": 102,
                    "right": 101
                },
                "oCloneOnDraw": {
                    "top": false,
                    "bottom": false,
                    "left": true,
                    "right": true
                },
                "oMes": {
                    "iTableWidth": 0,
                    "iTableHeight": 0,
                    "iTableLeft": 0,
                    "iTableRight": 0, /* note this is left+width, not actually "right" */
                    "iTableTop": 0,
                    "iTableBottom": 0 /* note this is top+height, not actually "bottom" */
                },
                "oOffset": {
                    "top": 0
                },
                "nTable": null,
                "bFooter": false,
                "bInitComplete": false
            };

            /*
             * Function: fnGetSettings
             * Purpose:  Get the settings for this object
             * Returns:  object: - settings object
             * Inputs:   -
             */
            this.fnGetSettings = function () {
                return oSettings;
            };

            /*
             * Function: fnUpdate
             * Purpose:  Update the positioning and copies of the fixed elements
             * Returns:  -
             * Inputs:   -
             */
            this.fnUpdate = function () {
                this._fnUpdateClones();
                this._fnUpdatePositions();
            };

            /*
             * Function: fnPosition
             * Purpose:  Update the positioning of the fixed elements
             * Returns:  -
             * Inputs:   -
             */
            this.fnPosition = function () {
                this._fnUpdatePositions();
            };


            var dt = $.fn.dataTable.Api ?
                new $.fn.dataTable.Api(mTable).settings()[0] :
                mTable.fnSettings();

            dt._oPluginFixedHeader = this;

            /* Let's do it */
            this.fnInit(dt, oInit);

        };


        /*
         * Variable: FixedHeader
         * Purpose:  Prototype for FixedHeader
         * Scope:    global
         */
        FixedHeader.prototype = {
            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Initialisation
             */

            /*
             * Function: fnInit
             * Purpose:  The "constructor"
             * Returns:  -
             * Inputs:   {as FixedHeader function}
             */
            fnInit: function (oDtSettings, oInit) {
                var s = this.fnGetSettings();
                var that = this;

                /* Record the user definable settings */
                this.fnInitSettings(s, oInit);

                if (oDtSettings.oScroll.sX !== "" || oDtSettings.oScroll.sY !== "") {
                    alert("FixedHeader 2 is not supported with DataTables' scrolling mode at this time");
                    return;
                }

                s.nTable = oDtSettings.nTable;
                oDtSettings.aoDrawCallback.unshift({
                    "fn": function () {
                        FixedHeader.fnMeasure();
                        that._fnUpdateClones.call(that);
                        that._fnUpdatePositions.call(that);
                    },
                    "sName": "FixedHeader"
                });

                s.bFooter = ($('>tfoot', s.nTable).length > 0) ? true : false;

                /* Add the 'sides' that are fixed */
                if (s.oSides.top) {
                    s.aoCache.push(that._fnCloneTable("fixedHeader", "FixedHeader_Header", that._fnCloneThead));
                }
                if (s.oSides.bottom) {
                    s.aoCache.push(that._fnCloneTable("fixedFooter", "FixedHeader_Footer", that._fnCloneTfoot));
                }
                if (s.oSides.left) {
                    s.aoCache.push(that._fnCloneTable("fixedLeft", "FixedHeader_Left", that._fnCloneTLeft, s.oSides.left));
                }
                if (s.oSides.right) {
                    s.aoCache.push(that._fnCloneTable("fixedRight", "FixedHeader_Right", that._fnCloneTRight, s.oSides.right));
                }

                /* Event listeners for window movement */
                FixedHeader.afnScroll.push(function () {
                    that._fnUpdatePositions.call(that);
                });

                $(window).resize(function () {
                    FixedHeader.fnMeasure();
                    that._fnUpdateClones.call(that);
                    that._fnUpdatePositions.call(that);
                });

                $(s.nTable)
                    .on('column-reorder.dt', function () {
                        FixedHeader.fnMeasure();
                        that._fnUpdateClones(true);
                        that._fnUpdatePositions();
                    })
                    .on('column-visibility.dt', function () {
                        FixedHeader.fnMeasure();
                        that._fnUpdateClones(true);
                        that._fnUpdatePositions();
                    });

                /* Get things right to start with */
                FixedHeader.fnMeasure();
                that._fnUpdateClones();
                that._fnUpdatePositions();

                s.bInitComplete = true;
            },


            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Support functions
             */

            /*
             * Function: fnInitSettings
             * Purpose:  Take the user's settings and copy them to our local store
             * Returns:  -
             * Inputs:   object:s - the local settings object
             *           object:oInit - the user's settings object
             */
            fnInitSettings: function (s, oInit) {
                if (oInit !== undefined) {
                    if (oInit.top !== undefined) {
                        s.oSides.top = oInit.top;
                    }
                    if (oInit.bottom !== undefined) {
                        s.oSides.bottom = oInit.bottom;
                    }
                    if (typeof oInit.left == 'boolean') {
                        s.oSides.left = oInit.left ? 1 : 0;
                    }
                    else if (oInit.left !== undefined) {
                        s.oSides.left = oInit.left;
                    }
                    if (typeof oInit.right == 'boolean') {
                        s.oSides.right = oInit.right ? 1 : 0;
                    }
                    else if (oInit.right !== undefined) {
                        s.oSides.right = oInit.right;
                    }

                    if (oInit.zTop !== undefined) {
                        s.oZIndexes.top = oInit.zTop;
                    }
                    if (oInit.zBottom !== undefined) {
                        s.oZIndexes.bottom = oInit.zBottom;
                    }
                    if (oInit.zLeft !== undefined) {
                        s.oZIndexes.left = oInit.zLeft;
                    }
                    if (oInit.zRight !== undefined) {
                        s.oZIndexes.right = oInit.zRight;
                    }

                    if (oInit.offsetTop !== undefined) {
                        s.oOffset.top = oInit.offsetTop;
                    }
                    if (oInit.alwaysCloneTop !== undefined) {
                        s.oCloneOnDraw.top = oInit.alwaysCloneTop;
                    }
                    if (oInit.alwaysCloneBottom !== undefined) {
                        s.oCloneOnDraw.bottom = oInit.alwaysCloneBottom;
                    }
                    if (oInit.alwaysCloneLeft !== undefined) {
                        s.oCloneOnDraw.left = oInit.alwaysCloneLeft;
                    }
                    if (oInit.alwaysCloneRight !== undefined) {
                        s.oCloneOnDraw.right = oInit.alwaysCloneRight;
                    }
                }
            },

            /*
             * Function: _fnCloneTable
             * Purpose:  Clone the table node and do basic initialisation
             * Returns:  -
             * Inputs:   -
             */
            _fnCloneTable: function (sType, sClass, fnClone, iCells) {
                var s = this.fnGetSettings();
                var nCTable;

                /* We know that the table _MUST_ has a DIV wrapped around it, because this is simply how
                 * DataTables works. Therefore, we can set this to be relatively position (if it is not
                 * alreadu absolute, and use this as the base point for the cloned header
                 */
                if ($(s.nTable.parentNode).css('position') != "absolute") {
                    s.nTable.parentNode.style.position = "relative";
                }

                /* Just a shallow clone will do - we only want the table node */
                nCTable = s.nTable.cloneNode(false);
                nCTable.removeAttribute('id');

                var nDiv = document.createElement('div');
                nDiv.style.position = "absolute";
                nDiv.style.top = "0px";
                nDiv.style.left = "0px";
                nDiv.className += " FixedHeader_Cloned " + sType + " " + sClass;

                /* Set the zIndexes */
                if (sType == "fixedHeader") {
                    nDiv.style.zIndex = s.oZIndexes.top;
                }
                if (sType == "fixedFooter") {
                    nDiv.style.zIndex = s.oZIndexes.bottom;
                }
                if (sType == "fixedLeft") {
                    nDiv.style.zIndex = s.oZIndexes.left;
                }
                else if (sType == "fixedRight") {
                    nDiv.style.zIndex = s.oZIndexes.right;
                }

                /* remove margins since we are going to position it absolutely */
                nCTable.style.margin = "0";

                /* Insert the newly cloned table into the DOM, on top of the "real" header */
                nDiv.appendChild(nCTable);
                document.body.appendChild(nDiv);

                return {
                    "nNode": nCTable,
                    "nWrapper": nDiv,
                    "sType": sType,
                    "sPosition": "",
                    "sTop": "",
                    "sLeft": "",
                    "fnClone": fnClone,
                    "iCells": iCells
                };
            },

            /*
             * Function: _fnMeasure
             * Purpose:  Get the current positioning of the table in the DOM
             * Returns:  -
             * Inputs:   -
             */
            _fnMeasure: function () {
                var
                    s = this.fnGetSettings(),
                    m = s.oMes,
                    jqTable = $(s.nTable),
                    oOffset = jqTable.offset(),
                    iParentScrollTop = this._fnSumScroll(s.nTable.parentNode, 'scrollTop'),
                    iParentScrollLeft = this._fnSumScroll(s.nTable.parentNode, 'scrollLeft');

                m.iTableWidth = jqTable.outerWidth();
                m.iTableHeight = jqTable.outerHeight();
                m.iTableLeft = oOffset.left + s.nTable.parentNode.scrollLeft;
                m.iTableTop = oOffset.top + iParentScrollTop;
                m.iTableRight = m.iTableLeft + m.iTableWidth;
                m.iTableRight = FixedHeader.oDoc.iWidth - m.iTableLeft - m.iTableWidth;
                m.iTableBottom = FixedHeader.oDoc.iHeight - m.iTableTop - m.iTableHeight;
            },

            /*
             * Function: _fnSumScroll
             * Purpose:  Sum node parameters all the way to the top
             * Returns:  int: sum
             * Inputs:   node:n - node to consider
             *           string:side - scrollTop or scrollLeft
             */
            _fnSumScroll: function (n, side) {
                var i = n[side];
                while (n = n.parentNode) {
                    if (n.nodeName == 'HTML' || n.nodeName == 'BODY') {
                        break;
                    }
                    i = n[side];
                }
                return i;
            },

            /*
             * Function: _fnUpdatePositions
             * Purpose:  Loop over the fixed elements for this table and update their positions
             * Returns:  -
             * Inputs:   -
             */
            _fnUpdatePositions: function () {
                var s = this.fnGetSettings();
                this._fnMeasure();

                for (var i = 0, iLen = s.aoCache.length ; i < iLen ; i++) {
                    if (s.aoCache[i].sType == "fixedHeader") {
                        this._fnScrollFixedHeader(s.aoCache[i]);
                    }
                    else if (s.aoCache[i].sType == "fixedFooter") {
                        this._fnScrollFixedFooter(s.aoCache[i]);
                    }
                    else if (s.aoCache[i].sType == "fixedLeft") {
                        this._fnScrollHorizontalLeft(s.aoCache[i]);
                    }
                    else {
                        this._fnScrollHorizontalRight(s.aoCache[i]);
                    }
                }
            },

            /*
             * Function: _fnUpdateClones
             * Purpose:  Loop over the fixed elements for this table and call their cloning functions
             * Returns:  -
             * Inputs:   -
             */
            _fnUpdateClones: function (full) {
                var s = this.fnGetSettings();

                if (full) {
                    // This is a little bit of a hack to force a full clone draw. When
                    // `full` is set to true, we want to reclone the source elements,
                    // regardless of the clone-on-draw settings
                    s.bInitComplete = false;
                }

                for (var i = 0, iLen = s.aoCache.length ; i < iLen ; i++) {
                    s.aoCache[i].fnClone.call(this, s.aoCache[i]);
                }

                if (full) {
                    s.bInitComplete = true;
                }
            },


            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Scrolling functions
             */

            /*
             * Function: _fnScrollHorizontalLeft
             * Purpose:  Update the positioning of the scrolling elements
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnScrollHorizontalRight: function (oCache) {
                var
                    s = this.fnGetSettings(),
                    oMes = s.oMes,
                    oWin = FixedHeader.oWin,
                    oDoc = FixedHeader.oDoc,
                    nTable = oCache.nWrapper,
                    iFixedWidth = $(nTable).outerWidth();

                if (oWin.iScrollRight < oMes.iTableRight) {
                    /* Fully right aligned */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', oMes.iTableTop + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', (oMes.iTableLeft + oMes.iTableWidth - iFixedWidth) + "px", 'left', nTable.style);
                }
                else if (oMes.iTableLeft < oDoc.iWidth - oWin.iScrollRight - iFixedWidth) {
                    /* Middle */
                    this._fnUpdateCache(oCache, 'sPosition', 'fixed', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oMes.iTableTop - oWin.iScrollTop) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', (oWin.iWidth - iFixedWidth) + "px", 'left', nTable.style);
                }
                else {
                    /* Fully left aligned */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', oMes.iTableTop + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
            },

            /*
             * Function: _fnScrollHorizontalLeft
             * Purpose:  Update the positioning of the scrolling elements
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnScrollHorizontalLeft: function (oCache) {
                var
                    s = this.fnGetSettings(),
                    oMes = s.oMes,
                    oWin = FixedHeader.oWin,
                    oDoc = FixedHeader.oDoc,
                    nTable = oCache.nWrapper,
                    iCellWidth = $(nTable).outerWidth();

                if (oWin.iScrollLeft < oMes.iTableLeft) {
                    /* Fully left align */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', oMes.iTableTop + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
                else if (oWin.iScrollLeft < oMes.iTableLeft + oMes.iTableWidth - iCellWidth) {
                    this._fnUpdateCache(oCache, 'sPosition', 'fixed', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oMes.iTableTop - oWin.iScrollTop) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', "0px", 'left', nTable.style);
                }
                else {
                    /* Fully right align */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', oMes.iTableTop + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', (oMes.iTableLeft + oMes.iTableWidth - iCellWidth) + "px", 'left', nTable.style);
                }
            },

            /*
             * Function: _fnScrollFixedFooter
             * Purpose:  Update the positioning of the scrolling elements
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnScrollFixedFooter: function (oCache) {
                var
                    s = this.fnGetSettings(),
                    oMes = s.oMes,
                    oWin = FixedHeader.oWin,
                    oDoc = FixedHeader.oDoc,
                    nTable = oCache.nWrapper,
                    iTheadHeight = $("thead", s.nTable).outerHeight(),
                    iCellHeight = $(nTable).outerHeight();

                if (oWin.iScrollBottom < oMes.iTableBottom) {
                    /* Below */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oMes.iTableTop + oMes.iTableHeight - iCellHeight) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
                else if (oWin.iScrollBottom < oMes.iTableBottom + oMes.iTableHeight - iCellHeight - iTheadHeight) {
                    this._fnUpdateCache(oCache, 'sPosition', 'fixed', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oWin.iHeight - iCellHeight) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', (oMes.iTableLeft - oWin.iScrollLeft) + "px", 'left', nTable.style);
                }
                else {
                    /* Above */
                    this._fnUpdateCache(oCache, 'sPosition', 'absolute', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oMes.iTableTop + iCellHeight) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
            },

            /*
             * Function: _fnScrollFixedHeader
             * Purpose:  Update the positioning of the scrolling elements
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnScrollFixedHeader: function (oCache) {
                var
                    s = this.fnGetSettings(),
                    oMes = s.oMes,
                    oWin = FixedHeader.oWin,
                    oDoc = FixedHeader.oDoc,
                    nTable = oCache.nWrapper,
                    iTbodyHeight = 0,
                    anTbodies = s.nTable.getElementsByTagName('tbody');

                for (var i = 0; i < anTbodies.length; ++i) {
                    iTbodyHeight += anTbodies[i].offsetHeight;
                }

                if (oMes.iTableTop > oWin.iScrollTop + s.oOffset.top) {
                    /* Above the table */
                    this._fnUpdateCache(oCache, 'sPosition', "absolute", 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', oMes.iTableTop + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
                else if (oWin.iScrollTop + s.oOffset.top > oMes.iTableTop + iTbodyHeight) {
                    /* At the bottom of the table */
                    this._fnUpdateCache(oCache, 'sPosition', "absolute", 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', (oMes.iTableTop + iTbodyHeight) + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', oMes.iTableLeft + "px", 'left', nTable.style);
                }
                else {
                    /* In the middle of the table */
                    this._fnUpdateCache(oCache, 'sPosition', 'fixed', 'position', nTable.style);
                    this._fnUpdateCache(oCache, 'sTop', s.oOffset.top + "px", 'top', nTable.style);
                    this._fnUpdateCache(oCache, 'sLeft', (oMes.iTableLeft - oWin.iScrollLeft) + "px", 'left', nTable.style);
                }
            },

            /*
             * Function: _fnUpdateCache
             * Purpose:  Check the cache and update cache and value if needed
             * Returns:  -
             * Inputs:   object:oCache - local cache object
             *           string:sCache - cache property
             *           string:sSet - value to set
             *           string:sProperty - object property to set
             *           object:oObj - object to update
             */
            _fnUpdateCache: function (oCache, sCache, sSet, sProperty, oObj) {
                if (oCache[sCache] != sSet) {
                    oObj[sProperty] = sSet;
                    oCache[sCache] = sSet;
                }
            },



            /**
             * Copy the classes of all child nodes from one element to another. This implies
             * that the two have identical structure - no error checking is performed to that
             * fact.
             *  @param {element} source Node to copy classes from
             *  @param {element} dest Node to copy classes too
             */
            _fnClassUpdate: function (source, dest) {
                var that = this;

                if (source.nodeName.toUpperCase() === "TR" || source.nodeName.toUpperCase() === "TH" ||
                     source.nodeName.toUpperCase() === "TD" || source.nodeName.toUpperCase() === "SPAN") {
                    dest.className = source.className;
                }

                $(source).children().each(function (i) {
                    that._fnClassUpdate($(source).children()[i], $(dest).children()[i]);
                });
            },


            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Cloning functions
             */

            /*
             * Function: _fnCloneThead
             * Purpose:  Clone the thead element
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnCloneThead: function (oCache) {
                var s = this.fnGetSettings();
                var nTable = oCache.nNode;

                if (s.bInitComplete && !s.oCloneOnDraw.top) {
                    this._fnClassUpdate($('thead', s.nTable)[0], $('thead', nTable)[0]);
                    return;
                }

                /* Set the wrapper width to match that of the cloned table */
                var iDtWidth = $(s.nTable).outerWidth();
                oCache.nWrapper.style.width = iDtWidth + "px";
                nTable.style.width = iDtWidth + "px";

                /* Remove any children the cloned table has */
                while (nTable.childNodes.length > 0) {
                    $('thead th', nTable).unbind('click');
                    nTable.removeChild(nTable.childNodes[0]);
                }

                /* Clone the DataTables header */
                var nThead = $('thead', s.nTable).clone(true)[0];
                nTable.appendChild(nThead);

                /* Copy the widths across - apparently a clone isn't good enough for this */
                var a = [];
                var b = [];

                $("thead>tr th", s.nTable).each(function (i) {
                    a.push($(this).width());
                });

                $("thead>tr td", s.nTable).each(function (i) {
                    b.push($(this).width());
                });

                $("thead>tr th", s.nTable).each(function (i) {
                    $("thead>tr th:eq(" + i + ")", nTable).width(a[i]);
                    $(this).width(a[i]);
                });

                $("thead>tr td", s.nTable).each(function (i) {
                    $("thead>tr td:eq(" + i + ")", nTable).width(b[i]);
                    $(this).width(b[i]);
                });

                // Stop DataTables 1.9 from putting a focus ring on the headers when
                // clicked to sort
                $('th.sorting, th.sorting_desc, th.sorting_asc', nTable).bind('click', function () {
                    this.blur();
                });
            },

            /*
             * Function: _fnCloneTfoot
             * Purpose:  Clone the tfoot element
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnCloneTfoot: function (oCache) {
                var s = this.fnGetSettings();
                var nTable = oCache.nNode;

                /* Set the wrapper width to match that of the cloned table */
                oCache.nWrapper.style.width = $(s.nTable).outerWidth() + "px";

                /* Remove any children the cloned table has */
                while (nTable.childNodes.length > 0) {
                    nTable.removeChild(nTable.childNodes[0]);
                }

                /* Clone the DataTables footer */
                var nTfoot = $('tfoot', s.nTable).clone(true)[0];
                nTable.appendChild(nTfoot);

                /* Copy the widths across - apparently a clone isn't good enough for this */
                $("tfoot:eq(0)>tr th", s.nTable).each(function (i) {
                    $("tfoot:eq(0)>tr th:eq(" + i + ")", nTable).width($(this).width());
                });

                $("tfoot:eq(0)>tr td", s.nTable).each(function (i) {
                    $("tfoot:eq(0)>tr td:eq(" + i + ")", nTable).width($(this).width());
                });
            },

            /*
             * Function: _fnCloneTLeft
             * Purpose:  Clone the left column(s)
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnCloneTLeft: function (oCache) {
                var s = this.fnGetSettings();
                var nTable = oCache.nNode;
                var nBody = $('tbody', s.nTable)[0];

                /* Remove any children the cloned table has */
                while (nTable.childNodes.length > 0) {
                    nTable.removeChild(nTable.childNodes[0]);
                }

                /* Is this the most efficient way to do this - it looks horrible... */
                nTable.appendChild($("thead", s.nTable).clone(true)[0]);
                nTable.appendChild($("tbody", s.nTable).clone(true)[0]);
                if (s.bFooter) {
                    nTable.appendChild($("tfoot", s.nTable).clone(true)[0]);
                }

                /* Remove unneeded cells */
                var sSelector = 'gt(' + (oCache.iCells - 1) + ')';
                $('thead tr', nTable).each(function (k) {
                    $('th:' + sSelector, this).remove();
                });

                $('tfoot tr', nTable).each(function (k) {
                    $('th:' + sSelector, this).remove();
                });

                $('tbody tr', nTable).each(function (k) {
                    $('td:' + sSelector, this).remove();
                });

                this.fnEqualiseHeights('thead', nBody.parentNode, nTable);
                this.fnEqualiseHeights('tbody', nBody.parentNode, nTable);
                this.fnEqualiseHeights('tfoot', nBody.parentNode, nTable);

                var iWidth = 0;
                for (var i = 0; i < oCache.iCells; i++) {
                    iWidth += $('thead tr th:eq(' + i + ')', s.nTable).outerWidth();
                }
                nTable.style.width = iWidth + "px";
                oCache.nWrapper.style.width = iWidth + "px";
            },

            /*
             * Function: _fnCloneTRight
             * Purpose:  Clone the right most column(s)
             * Returns:  -
             * Inputs:   object:oCache - the cached values for this fixed element
             */
            _fnCloneTRight: function (oCache) {
                var s = this.fnGetSettings();
                var nBody = $('tbody', s.nTable)[0];
                var nTable = oCache.nNode;
                var iCols = $('tbody tr:eq(0) td', s.nTable).length;

                /* Remove any children the cloned table has */
                while (nTable.childNodes.length > 0) {
                    nTable.removeChild(nTable.childNodes[0]);
                }

                /* Is this the most efficient way to do this - it looks horrible... */
                nTable.appendChild($("thead", s.nTable).clone(true)[0]);
                nTable.appendChild($("tbody", s.nTable).clone(true)[0]);
                if (s.bFooter) {
                    nTable.appendChild($("tfoot", s.nTable).clone(true)[0]);
                }
                $('thead tr th:lt(' + (iCols - oCache.iCells) + ')', nTable).remove();
                $('tfoot tr th:lt(' + (iCols - oCache.iCells) + ')', nTable).remove();

                /* Remove unneeded cells */
                $('tbody tr', nTable).each(function (k) {
                    $('td:lt(' + (iCols - oCache.iCells) + ')', this).remove();
                });

                this.fnEqualiseHeights('thead', nBody.parentNode, nTable);
                this.fnEqualiseHeights('tbody', nBody.parentNode, nTable);
                this.fnEqualiseHeights('tfoot', nBody.parentNode, nTable);

                var iWidth = 0;
                for (var i = 0; i < oCache.iCells; i++) {
                    iWidth += $('thead tr th:eq(' + (iCols - 1 - i) + ')', s.nTable).outerWidth();
                }
                nTable.style.width = iWidth + "px";
                oCache.nWrapper.style.width = iWidth + "px";
            },


            /**
             * Equalise the heights of the rows in a given table node in a cross browser way. Note that this
             * is more or less lifted as is from FixedColumns
             *  @method  fnEqualiseHeights
             *  @returns void
             *  @param   {string} parent Node type - thead, tbody or tfoot
             *  @param   {element} original Original node to take the heights from
             *  @param   {element} clone Copy the heights to
             *  @private
             */
            "fnEqualiseHeights": function (parent, original, clone) {
                var that = this;
                var originals = $(parent + ' tr', original);
                var height;

                $(parent + ' tr', clone).each(function (k) {
                    height = originals.eq(k).css('height');

                    // This is nasty :-(. IE has a sub-pixel error even when setting
                    // the height below (the Firefox fix) which causes the fixed column
                    // to go out of alignment. Need to add a pixel before the assignment
                    // Can this be feature detected? Not sure how...
                    if (navigator.appName == 'Microsoft Internet Explorer') {
                        height = parseInt(height, 10) + 1;
                    }

                    $(this).css('height', height);

                    // For Firefox to work, we need to also set the height of the
                    // original row, to the value that we read from it! Otherwise there
                    // is a sub-pixel rounding error
                    originals.eq(k).css('height', height);
                });
            }
        };


        /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
         * Static properties and methods
         *   We use these for speed! This information is common to all instances of FixedHeader, so no
         * point if having them calculated and stored for each different instance.
         */

        /*
         * Variable: oWin
         * Purpose:  Store information about the window positioning
         * Scope:    FixedHeader
         */
        FixedHeader.oWin = {
            "iScrollTop": 0,
            "iScrollRight": 0,
            "iScrollBottom": 0,
            "iScrollLeft": 0,
            "iHeight": 0,
            "iWidth": 0
        };

        /*
         * Variable: oDoc
         * Purpose:  Store information about the document size
         * Scope:    FixedHeader
         */
        FixedHeader.oDoc = {
            "iHeight": 0,
            "iWidth": 0
        };

        /*
         * Variable: afnScroll
         * Purpose:  Array of functions that are to be used for the scrolling components
         * Scope:    FixedHeader
         */
        FixedHeader.afnScroll = [];

        /*
         * Function: fnMeasure
         * Purpose:  Update the measurements for the window and document
         * Returns:  -
         * Inputs:   -
         */
        FixedHeader.fnMeasure = function () {
            var
                jqWin = $(window),
                jqDoc = $(document),
                oWin = FixedHeader.oWin,
                oDoc = FixedHeader.oDoc;

            oDoc.iHeight = jqDoc.height();
            oDoc.iWidth = jqDoc.width();

            oWin.iHeight = jqWin.height();
            oWin.iWidth = jqWin.width();
            oWin.iScrollTop = jqWin.scrollTop();
            oWin.iScrollLeft = jqWin.scrollLeft();
            oWin.iScrollRight = oDoc.iWidth - oWin.iScrollLeft - oWin.iWidth;
            oWin.iScrollBottom = oDoc.iHeight - oWin.iScrollTop - oWin.iHeight;
        };


        FixedHeader.version = "2.1.2";


        /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
         * Global processing
         */

        /*
         * Just one 'scroll' event handler in FixedHeader, which calls the required components. This is
         * done as an optimisation, to reduce calculation and proagation time
         */
        $(window).scroll(function () {
            FixedHeader.fnMeasure();

            for (var i = 0, iLen = FixedHeader.afnScroll.length ; i < iLen ; i++) {
                FixedHeader.afnScroll[i]();
            }
        });


        $.fn.dataTable.FixedHeader = FixedHeader;
        $.fn.DataTable.FixedHeader = FixedHeader;


        return FixedHeader;
    }; // /factory


    // Define as an AMD module if possible
    if (typeof define === 'function' && define.amd) {
        define(['jquery.ba-resize', 'datatables'], factory);
    }
    else if (typeof exports === 'object') {
        // Node/CommonJS
        factory(require('jquery'), require('datatables'));
    }
    else if (jQuery && !jQuery.fn.dataTable.FixedHeader) {
        // Otherwise simply initialise as normal, stopping multiple evaluation
        factory(jQuery, jQuery.fn.dataTable);
    }


})(window, document);

﻿/*! FixedColumns 3.0.4
 * 漏2010-2014 SpryMedia Ltd - datatables.net/license
 */

/**
 * @summary     FixedColumns
 * @description Freeze columns in place on a scrolling DataTable
 * @version     3.0.4
 * @file        dataTables.fixedColumns.js
 * @author      SpryMedia Ltd (www.sprymedia.co.uk)
 * @contact     www.sprymedia.co.uk/contact
 * @copyright   Copyright 2010-2014 SpryMedia Ltd.
 *
 * This source file is free software, available under the following license:
 *   MIT license - http://datatables.net/license/mit
 *
 * This source file is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE. See the license files for details.
 *
 * For details please refer to: http://www.datatables.net
 */


(function (window, document, undefined) {


    var factory = function ($, DataTable) {
        "use strict";

        /**
         * When making use of DataTables' x-axis scrolling feature, you may wish to
         * fix the left most column in place. This plug-in for DataTables provides
         * exactly this option (note for non-scrolling tables, please use the
         * FixedHeader plug-in, which can fix headers, footers and columns). Key
         * features include:
         *
         * * Freezes the left or right most columns to the side of the table
         * * Option to freeze two or more columns
         * * Full integration with DataTables' scrolling options
         * * Speed - FixedColumns is fast in its operation
         *
         *  @class
         *  @constructor
         *  @global
         *  @param {object} dt DataTables instance. With DataTables 1.10 this can also
         *    be a jQuery collection, a jQuery selector, DataTables API instance or
         *    settings object.
         *  @param {object} [init={}] Configuration object for FixedColumns. Options are
         *    defined by {@link FixedColumns.defaults}
         *
         *  @requires jQuery 1.7+
         *  @requires DataTables 1.8.0+
         *
         *  @example
         *      var table = $('#example').dataTable( {
         *        "scrollX": "100%"
         *      } );
         *      new $.fn.dataTable.fixedColumns( table );
         */
        var FixedColumns = function (dt, init) {
            var that = this;

            /* Sanity check - you just know it will happen */
            if (!(this instanceof FixedColumns)) {
                alert("FixedColumns warning: FixedColumns must be initialised with the 'new' keyword.");
                return;
            }

            if (typeof init == 'undefined') {
                init = {};
            }

            // Use the DataTables Hungarian notation mapping method, if it exists to
            // provide forwards compatibility for camel case variables
            var camelToHungarian = $.fn.dataTable.camelToHungarian;
            if (camelToHungarian) {
                camelToHungarian(FixedColumns.defaults, FixedColumns.defaults, true);
                camelToHungarian(FixedColumns.defaults, init);
            }

            // v1.10 allows the settings object to be got form a number of sources
            var dtSettings = $.fn.dataTable.Api ?
                new $.fn.dataTable.Api(dt).settings()[0] :
                dt.fnSettings();

            /**
             * Settings object which contains customisable information for FixedColumns instance
             * @namespace
             * @extends FixedColumns.defaults
             * @private
             */
            this.s = {
                /**
                 * DataTables settings objects
                 *  @type     object
                 *  @default  Obtained from DataTables instance
                 */
                "dt": dtSettings,

                /**
                 * Number of columns in the DataTable - stored for quick access
                 *  @type     int
                 *  @default  Obtained from DataTables instance
                 */
                "iTableColumns": dtSettings.aoColumns.length,

                /**
                 * Original outer widths of the columns as rendered by DataTables - used to calculate
                 * the FixedColumns grid bounding box
                 *  @type     array.<int>
                 *  @default  []
                 */
                "aiOuterWidths": [],

                /**
                 * Original inner widths of the columns as rendered by DataTables - used to apply widths
                 * to the columns
                 *  @type     array.<int>
                 *  @default  []
                 */
                "aiInnerWidths": []
            };


            /**
             * DOM elements used by the class instance
             * @namespace
             * @private
             *
             */
            this.dom = {
                /**
                 * DataTables scrolling element
                 *  @type     node
                 *  @default  null
                 */
                "scroller": null,

                /**
                 * DataTables header table
                 *  @type     node
                 *  @default  null
                 */
                "header": null,

                /**
                 * DataTables body table
                 *  @type     node
                 *  @default  null
                 */
                "body": null,

                /**
                 * DataTables footer table
                 *  @type     node
                 *  @default  null
                 */
                "footer": null,

                /**
                 * Display grid elements
                 * @namespace
                 */
                "grid": {
                    /**
                     * Grid wrapper. This is the container element for the 3x3 grid
                     *  @type     node
                     *  @default  null
                     */
                    "wrapper": null,

                    /**
                     * DataTables scrolling element. This element is the DataTables
                     * component in the display grid (making up the main table - i.e.
                     * not the fixed columns).
                     *  @type     node
                     *  @default  null
                     */
                    "dt": null,

                    /**
                     * Left fixed column grid components
                     * @namespace
                     */
                    "left": {
                        "wrapper": null,
                        "head": null,
                        "body": null,
                        "foot": null
                    },

                    /**
                     * Right fixed column grid components
                     * @namespace
                     */
                    "right": {
                        "wrapper": null,
                        "head": null,
                        "body": null,
                        "foot": null
                    }
                },

                /**
                 * Cloned table nodes
                 * @namespace
                 */
                "clone": {
                    /**
                     * Left column cloned table nodes
                     * @namespace
                     */
                    "left": {
                        /**
                         * Cloned header table
                         *  @type     node
                         *  @default  null
                         */
                        "header": null,

                        /**
                         * Cloned body table
                         *  @type     node
                         *  @default  null
                         */
                        "body": null,

                        /**
                         * Cloned footer table
                         *  @type     node
                         *  @default  null
                         */
                        "footer": null
                    },

                    /**
                     * Right column cloned table nodes
                     * @namespace
                     */
                    "right": {
                        /**
                         * Cloned header table
                         *  @type     node
                         *  @default  null
                         */
                        "header": null,

                        /**
                         * Cloned body table
                         *  @type     node
                         *  @default  null
                         */
                        "body": null,

                        /**
                         * Cloned footer table
                         *  @type     node
                         *  @default  null
                         */
                        "footer": null
                    }
                }
            };

            /* Attach the instance to the DataTables instance so it can be accessed easily */
            dtSettings._oFixedColumns = this;

            /* Let's do it */
            if (!dtSettings._bInitComplete) {
                dtSettings.oApi._fnCallbackReg(dtSettings, 'aoInitComplete', function () {
                    that._fnConstruct(init);
                }, 'FixedColumns');
            }
            else {
                this._fnConstruct(init);
            }
        };



        FixedColumns.prototype = /** @lends FixedColumns.prototype */{
            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Public methods
             * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

            /**
             * Update the fixed columns - including headers and footers. Note that FixedColumns will
             * automatically update the display whenever the host DataTable redraws.
             *  @returns {void}
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      var fc = new $.fn.dataTable.fixedColumns( table );
             *
             *      // at some later point when the table has been manipulated....
             *      fc.fnUpdate();
             */
            "fnUpdate": function () {
                this._fnDraw(true);
            },


            /**
             * Recalculate the resizes of the 3x3 grid that FixedColumns uses for display of the table.
             * This is useful if you update the width of the table container. Note that FixedColumns will
             * perform this function automatically when the window.resize event is fired.
             *  @returns {void}
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      var fc = new $.fn.dataTable.fixedColumns( table );
             *
             *      // Resize the table container and then have FixedColumns adjust its layout....
             *      $('#content').width( 1200 );
             *      fc.fnRedrawLayout();
             */
            "fnRedrawLayout": function () {
                this._fnColCalc();
                this._fnGridLayout();
                this.fnUpdate();
            },


            /**
             * Mark a row such that it's height should be recalculated when using 'semiauto' row
             * height matching. This function will have no effect when 'none' or 'auto' row height
             * matching is used.
             *  @param   {Node} nTr TR element that should have it's height recalculated
             *  @returns {void}
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      var fc = new $.fn.dataTable.fixedColumns( table );
             *
             *      // manipulate the table - mark the row as needing an update then update the table
             *      // this allows the redraw performed by DataTables fnUpdate to recalculate the row
             *      // height
             *      fc.fnRecalculateHeight();
             *      table.fnUpdate( $('#example tbody tr:eq(0)')[0], ["insert date", 1, 2, 3 ... ]);
             */
            "fnRecalculateHeight": function (nTr) {
                delete nTr._DTTC_iHeight;
                nTr.style.height = 'auto';
            },


            /**
             * Set the height of a given row - provides cross browser compatibility
             *  @param   {Node} nTarget TR element that should have it's height recalculated
             *  @param   {int} iHeight Height in pixels to set
             *  @returns {void}
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      var fc = new $.fn.dataTable.fixedColumns( table );
             *
             *      // You may want to do this after manipulating a row in the fixed column
             *      fc.fnSetRowHeight( $('#example tbody tr:eq(0)')[0], 50 );
             */
            "fnSetRowHeight": function (nTarget, iHeight) {
                nTarget.style.height = iHeight + "px";
            },


            /**
             * Get data index information about a row or cell in the table body.
             * This function is functionally identical to fnGetPosition in DataTables,
             * taking the same parameter (TH, TD or TR node) and returning exactly the
             * the same information (data index information). THe difference between
             * the two is that this method takes into account the fixed columns in the
             * table, so you can pass in nodes from the master table, or the cloned
             * tables and get the index position for the data in the main table.
             *  @param {node} node TR, TH or TD element to get the information about
             *  @returns {int} If nNode is given as a TR, then a single index is 
             *    returned, or if given as a cell, an array of [row index, column index
             *    (visible), column index (all)] is given.
             */
            "fnGetPosition": function (node) {
                var idx;
                var inst = this.s.dt.oInstance;

                if (!$(node).parents('.DTFC_Cloned').length) {
                    // Not in a cloned table
                    return inst.fnGetPosition(node);
                }
                else {
                    // Its in the cloned table, so need to look up position
                    if (node.nodeName.toLowerCase() === 'tr') {
                        idx = $(node).index();
                        return inst.fnGetPosition($('tr', this.s.dt.nTBody)[idx]);
                    }
                    else {
                        var colIdx = $(node).index();
                        idx = $(node.parentNode).index();
                        var row = inst.fnGetPosition($('tr', this.s.dt.nTBody)[idx]);

                        return [
                            row,
                            colIdx,
                            inst.oApi._fnVisibleToColumnIndex(this.s.dt, colIdx)
                        ];
                    }
                }
            },



            /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
             * Private methods (they are of course public in JS, but recommended as private)
             * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

            /**
             * Initialisation for FixedColumns
             *  @param   {Object} oInit User settings for initialisation
             *  @returns {void}
             *  @private
             */
            "_fnConstruct": function (oInit) {
                var i, iLen, iWidth,
                    that = this;

                /* Sanity checking */
                if (typeof this.s.dt.oInstance.fnVersionCheck != 'function' ||
                     this.s.dt.oInstance.fnVersionCheck('1.8.0') !== true) {
                    alert("FixedColumns " + FixedColumns.VERSION + " required DataTables 1.8.0 or later. " +
                        "Please upgrade your DataTables installation");
                    return;
                }

                if (this.s.dt.oScroll.sX === "") {
                    this.s.dt.oInstance.oApi._fnLog(this.s.dt, 1, "FixedColumns is not needed (no " +
                        "x-scrolling in DataTables enabled), so no action will be taken. Use 'FixedHeader' for " +
                        "column fixing when scrolling is not enabled");
                    return;
                }

                /* Apply the settings from the user / defaults */
                this.s = $.extend(true, this.s, FixedColumns.defaults, oInit);

                /* Set up the DOM as we need it and cache nodes */
                var classes = this.s.dt.oClasses;
                this.dom.grid.dt = $(this.s.dt.nTable).parents('div.' + classes.sScrollWrapper)[0];
                this.dom.scroller = $('div.' + classes.sScrollBody, this.dom.grid.dt)[0];

                /* Set up the DOM that we want for the fixed column layout grid */
                this._fnColCalc();
                this._fnGridSetup();

                /* Event handlers */
                var mouseController;

                // When the body is scrolled - scroll the left and right columns
                $(this.dom.scroller)
                    .on('mouseover.DTFC touchstart.DTFC', function () {
                        mouseController = 'main';
                    })
                    .on('scroll.DTFC', function () {
                        if (mouseController === 'main') {
                            if (that.s.iLeftColumns > 0) {
                                that.dom.grid.left.liner.scrollTop = that.dom.scroller.scrollTop;
                            }
                            if (that.s.iRightColumns > 0) {
                                that.dom.grid.right.liner.scrollTop = that.dom.scroller.scrollTop;
                            }
                        }
                    });

                var wheelType = 'onwheel' in document.createElement('div') ?
                    'wheel.DTFC' :
                    'mousewheel.DTFC';

                if (that.s.iLeftColumns > 0) {
                    // When scrolling the left column, scroll the body and right column
                    $(that.dom.grid.left.liner)
                        .on('mouseover.DTFC touchstart.DTFC', function () {
                            mouseController = 'left';
                        })
                        .on('scroll.DTFC', function () {
                            if (mouseController === 'left') {
                                that.dom.scroller.scrollTop = that.dom.grid.left.liner.scrollTop;
                                if (that.s.iRightColumns > 0) {
                                    that.dom.grid.right.liner.scrollTop = that.dom.grid.left.liner.scrollTop;
                                }
                            }
                        })
                        .on(wheelType, function (e) { // xxx update the destroy as well
                            // Pass horizontal scrolling through
                            var xDelta = e.type === 'wheel' ?
                                -e.originalEvent.deltaX :
                                e.originalEvent.wheelDeltaX;
                            that.dom.scroller.scrollLeft -= xDelta;
                        });
                }

                if (that.s.iRightColumns > 0) {
                    // When scrolling the right column, scroll the body and the left column
                    $(that.dom.grid.right.liner)
                        .on('mouseover.DTFC touchstart.DTFC', function () {
                            mouseController = 'right';
                        })
                        .on('scroll.DTFC', function () {
                            if (mouseController === 'right') {
                                that.dom.scroller.scrollTop = that.dom.grid.right.liner.scrollTop;
                                if (that.s.iLeftColumns > 0) {
                                    that.dom.grid.left.liner.scrollTop = that.dom.grid.right.liner.scrollTop;
                                }
                            }
                        })
                        .on(wheelType, function (e) {
                            // Pass horizontal scrolling through
                            var xDelta = e.type === 'wheel' ?
                                -e.originalEvent.deltaX :
                                e.originalEvent.wheelDeltaX;
                            that.dom.scroller.scrollLeft -= xDelta;
                        });
                }

                $(window).on('resize.DTFC', function () {
                    that._fnGridLayout.call(that);
                });

                var bFirstDraw = true;
                var jqTable = $(this.s.dt.nTable);

                jqTable
                    .on('draw.dt.DTFC', function () {
                        that._fnDraw.call(that, bFirstDraw);
                        bFirstDraw = false;
                    })
                    .on('column-sizing.dt.DTFC', function () {
                        that._fnColCalc();
                        that._fnGridLayout(that);
                    })
                    .on('column-visibility.dt.DTFC', function () {
                        that._fnColCalc();
                        that._fnGridLayout(that);
                        that._fnDraw(true);
                    })
                    .on('destroy.dt.DTFC', function () {
                        jqTable.off('column-sizing.dt.DTFC destroy.dt.DTFC draw.dt.DTFC');

                        $(that.dom.scroller).off('scroll.DTFC mouseover.DTFC');
                        $(window).off('resize.DTFC');

                        $(that.dom.grid.left.liner).off('scroll.DTFC mouseover.DTFC ' + wheelType);
                        $(that.dom.grid.left.wrapper).remove();

                        $(that.dom.grid.right.liner).off('scroll.DTFC mouseover.DTFC ' + wheelType);
                        $(that.dom.grid.right.wrapper).remove();
                    });

                /* Get things right to start with - note that due to adjusting the columns, there must be
                 * another redraw of the main table. It doesn't need to be a full redraw however.
                 */
                this._fnGridLayout();
                this.s.dt.oInstance.fnDraw(false);
            },


            /**
             * Calculate the column widths for the grid layout
             *  @returns {void}
             *  @private
             */
            "_fnColCalc": function () {
                var that = this;
                var iLeftWidth = 0;
                var iRightWidth = 0;

                this.s.aiInnerWidths = [];
                this.s.aiOuterWidths = [];

                $.each(this.s.dt.aoColumns, function (i, col) {
                    var th = $(col.nTh);
                    var border;

                    if (!th.filter(':visible').length) {
                        that.s.aiInnerWidths.push(0);
                        that.s.aiOuterWidths.push(0);
                    }
                    else {
                        // Inner width is used to assign widths to cells
                        // Outer width is used to calculate the container
                        var iWidth = th.outerWidth();

                        // When working with the left most-cell, need to add on the
                        // table's border to the outerWidth, since we need to take
                        // account of it, but it isn't in any cell
                        if (that.s.aiOuterWidths.length === 0) {
                            border = $(that.s.dt.nTable).css('border-left-width');
                            iWidth += typeof border === 'string' ? 1 : parseInt(border, 10);
                        }

                        // Likewise with the final column on the right
                        if (that.s.aiOuterWidths.length === that.s.dt.aoColumns.length - 1) {
                            border = $(that.s.dt.nTable).css('border-right-width');
                            iWidth += typeof border === 'string' ? 1 : parseInt(border, 10);
                        }

                        that.s.aiOuterWidths.push(iWidth);
                        that.s.aiInnerWidths.push(th.width());

                        if (i < that.s.iLeftColumns) {
                            iLeftWidth += iWidth;
                        }

                        if (that.s.iTableColumns - that.s.iRightColumns <= i) {
                            iRightWidth += iWidth;
                        }
                    }
                });

                this.s.iLeftWidth = iLeftWidth;
                this.s.iRightWidth = iRightWidth;
            },


            /**
             * Set up the DOM for the fixed column. The way the layout works is to create a 1x3 grid
             * for the left column, the DataTable (for which we just reuse the scrolling element DataTable
             * puts into the DOM) and the right column. In each of he two fixed column elements there is a
             * grouping wrapper element and then a head, body and footer wrapper. In each of these we then
             * place the cloned header, body or footer tables. This effectively gives as 3x3 grid structure.
             *  @returns {void}
             *  @private
             */
            "_fnGridSetup": function () {
                var that = this;
                var oOverflow = this._fnDTOverflow();
                var block;

                this.dom.body = this.s.dt.nTable;
                this.dom.header = this.s.dt.nTHead.parentNode;
                this.dom.header.parentNode.parentNode.style.position = "relative";

                var nSWrapper =
                    $('<div class="DTFC_ScrollWrapper" style="position:relative; clear:both;">' +
                        '<div class="DTFC_LeftWrapper" style="position:absolute; top:0; left:0;">' +
                            '<div class="DTFC_LeftHeadWrapper" style="position:relative; top:0; left:0; overflow:hidden;"></div>' +
                            '<div class="DTFC_LeftBodyWrapper" style="position:relative; top:0; left:0; overflow:hidden;">' +
                                '<div class="DTFC_LeftBodyLiner" style="position:relative; top:0; left:0; overflow-y:scroll;"></div>' +
                            '</div>' +
                            '<div class="DTFC_LeftFootWrapper" style="position:relative; top:0; left:0; overflow:hidden;"></div>' +
                        '</div>' +
                        '<div class="DTFC_RightWrapper" style="position:absolute; top:0; left:0;">' +
                            '<div class="DTFC_RightHeadWrapper" style="position:relative; top:0; left:0;">' +
                                '<div class="DTFC_RightHeadBlocker DTFC_Blocker" style="position:absolute; top:0; bottom:0;"></div>' +
                            '</div>' +
                            '<div class="DTFC_RightBodyWrapper" style="position:relative; top:0; left:0; overflow:hidden;">' +
                                '<div class="DTFC_RightBodyLiner" style="position:relative; top:0; left:0; overflow-y:scroll;"></div>' +
                            '</div>' +
                            '<div class="DTFC_RightFootWrapper" style="position:relative; top:0; left:0;">' +
                                '<div class="DTFC_RightFootBlocker DTFC_Blocker" style="position:absolute; top:0; bottom:0;"></div>' +
                            '</div>' +
                        '</div>' +
                    '</div>')[0];
                var nLeft = nSWrapper.childNodes[0];
                var nRight = nSWrapper.childNodes[1];

                this.dom.grid.dt.parentNode.insertBefore(nSWrapper, this.dom.grid.dt);
                nSWrapper.appendChild(this.dom.grid.dt);

                this.dom.grid.wrapper = nSWrapper;

                if (this.s.iLeftColumns > 0) {
                    this.dom.grid.left.wrapper = nLeft;
                    this.dom.grid.left.head = nLeft.childNodes[0];
                    this.dom.grid.left.body = nLeft.childNodes[1];
                    this.dom.grid.left.liner = $('div.DTFC_LeftBodyLiner', nSWrapper)[0];

                    nSWrapper.appendChild(nLeft);
                }

                if (this.s.iRightColumns > 0) {
                    this.dom.grid.right.wrapper = nRight;
                    this.dom.grid.right.head = nRight.childNodes[0];
                    this.dom.grid.right.body = nRight.childNodes[1];
                    this.dom.grid.right.liner = $('div.DTFC_RightBodyLiner', nSWrapper)[0];

                    block = $('div.DTFC_RightHeadBlocker', nSWrapper)[0];
                    block.style.width = oOverflow.bar + "px";
                    block.style.right = -oOverflow.bar + "px";
                    this.dom.grid.right.headBlock = block;

                    block = $('div.DTFC_RightFootBlocker', nSWrapper)[0];
                    block.style.width = oOverflow.bar + "px";
                    block.style.right = -oOverflow.bar + "px";
                    this.dom.grid.right.footBlock = block;

                    nSWrapper.appendChild(nRight);
                }

                if (this.s.dt.nTFoot) {
                    this.dom.footer = this.s.dt.nTFoot.parentNode;
                    if (this.s.iLeftColumns > 0) {
                        this.dom.grid.left.foot = nLeft.childNodes[2];
                    }
                    if (this.s.iRightColumns > 0) {
                        this.dom.grid.right.foot = nRight.childNodes[2];
                    }
                }
            },


            /**
             * Style and position the grid used for the FixedColumns layout
             *  @returns {void}
             *  @private
             */
            "_fnGridLayout": function () {
                var oGrid = this.dom.grid;
                var iWidth = $(oGrid.wrapper).width();
                var iBodyHeight = $(this.s.dt.nTable.parentNode).outerHeight();
                var iFullHeight = $(this.s.dt.nTable.parentNode.parentNode).outerHeight();
                var oOverflow = this._fnDTOverflow();
                var
                    iLeftWidth = this.s.iLeftWidth,
                    iRightWidth = this.s.iRightWidth,
                    iRight;
                var scrollbarAdjust = function (node, width) {
                    if (!oOverflow.bar) {
                        // If there is no scrollbar (Macs) we need to hide the auto scrollbar
                        node.style.width = (width + 20) + "px";
                        node.style.paddingRight = "20px";
                        node.style.boxSizing = "border-box";
                    }
                    else {
                        // Otherwise just overflow by the scrollbar
                        node.style.width = (width + oOverflow.bar) + "px";
                    }
                };

                // When x scrolling - don't paint the fixed columns over the x scrollbar
                if (oOverflow.x) {
                    iBodyHeight -= oOverflow.bar;
                }

                oGrid.wrapper.style.height = iFullHeight + "px";

                if (this.s.iLeftColumns > 0) {
                    oGrid.left.wrapper.style.width = iLeftWidth + "px";
                    oGrid.left.wrapper.style.height = "1px";
                    oGrid.left.body.style.height = iBodyHeight + "px";
                    if (oGrid.left.foot) {
                        oGrid.left.foot.style.top = (oOverflow.x ? oOverflow.bar : 0) + "px"; // shift footer for scrollbar
                    }

                    scrollbarAdjust(oGrid.left.liner, iLeftWidth);
                    oGrid.left.liner.style.height = iBodyHeight + "px";
                }

                if (this.s.iRightColumns > 0) {
                    iRight = iWidth - iRightWidth;
                    if (oOverflow.y) {
                        iRight -= oOverflow.bar;
                    }

                    oGrid.right.wrapper.style.width = iRightWidth + "px";
                    oGrid.right.wrapper.style.left = iRight + "px";
                    oGrid.right.wrapper.style.height = "1px";
                    oGrid.right.body.style.height = iBodyHeight + "px";
                    if (oGrid.right.foot) {
                        oGrid.right.foot.style.top = (oOverflow.x ? oOverflow.bar : 0) + "px";
                    }

                    scrollbarAdjust(oGrid.right.liner, iRightWidth);
                    oGrid.right.liner.style.height = iBodyHeight + "px";

                    oGrid.right.headBlock.style.display = oOverflow.y ? 'block' : 'none';
                    oGrid.right.footBlock.style.display = oOverflow.y ? 'block' : 'none';
                }
            },


            /**
             * Get information about the DataTable's scrolling state - specifically if the table is scrolling
             * on either the x or y axis, and also the scrollbar width.
             *  @returns {object} Information about the DataTables scrolling state with the properties:
             *    'x', 'y' and 'bar'
             *  @private
             */
            "_fnDTOverflow": function () {
                var nTable = this.s.dt.nTable;
                var nTableScrollBody = nTable.parentNode;
                var out = {
                    "x": false,
                    "y": false,
                    "bar": this.s.dt.oScroll.iBarWidth
                };

                if (nTable.offsetWidth > nTableScrollBody.clientWidth) {
                    out.x = true;
                }

                if (nTable.offsetHeight > nTableScrollBody.clientHeight) {
                    out.y = true;
                }

                return out;
            },


            /**
             * Clone and position the fixed columns
             *  @returns {void}
             *  @param   {Boolean} bAll Indicate if the header and footer should be updated as well (true)
             *  @private
             */
            "_fnDraw": function (bAll) {
                this._fnGridLayout();
                this._fnCloneLeft(bAll);
                this._fnCloneRight(bAll);

                /* Draw callback function */
                if (this.s.fnDrawCallback !== null) {
                    this.s.fnDrawCallback.call(this, this.dom.clone.left, this.dom.clone.right);
                }

                /* Event triggering */
                $(this).trigger('draw.dtfc', {
                    "leftClone": this.dom.clone.left,
                    "rightClone": this.dom.clone.right
                });
            },


            /**
             * Clone the right columns
             *  @returns {void}
             *  @param   {Boolean} bAll Indicate if the header and footer should be updated as well (true)
             *  @private
             */
            "_fnCloneRight": function (bAll) {
                if (this.s.iRightColumns <= 0) {
                    return;
                }

                var that = this,
                    i, jq,
                    aiColumns = [];

                for (i = this.s.iTableColumns - this.s.iRightColumns ; i < this.s.iTableColumns ; i++) {
                    if (this.s.dt.aoColumns[i].bVisible) {
                        aiColumns.push(i);
                    }
                }

                this._fnClone(this.dom.clone.right, this.dom.grid.right, aiColumns, bAll);
            },


            /**
             * Clone the left columns
             *  @returns {void}
             *  @param   {Boolean} bAll Indicate if the header and footer should be updated as well (true)
             *  @private
             */
            "_fnCloneLeft": function (bAll) {
                if (this.s.iLeftColumns <= 0) {
                    return;
                }

                var that = this,
                    i, jq,
                    aiColumns = [];

                for (i = 0 ; i < this.s.iLeftColumns ; i++) {
                    if (this.s.dt.aoColumns[i].bVisible) {
                        aiColumns.push(i);
                    }
                }

                this._fnClone(this.dom.clone.left, this.dom.grid.left, aiColumns, bAll);
            },


            /**
             * Make a copy of the layout object for a header or footer element from DataTables. Note that
             * this method will clone the nodes in the layout object.
             *  @returns {Array} Copy of the layout array
             *  @param   {Object} aoOriginal Layout array from DataTables (aoHeader or aoFooter)
             *  @param   {Object} aiColumns Columns to copy
             *  @private
             */
            "_fnCopyLayout": function (aoOriginal, aiColumns) {
                var aReturn = [];
                var aClones = [];
                var aCloned = [];

                for (var i = 0, iLen = aoOriginal.length ; i < iLen ; i++) {
                    var aRow = [];
                    aRow.nTr = $(aoOriginal[i].nTr).clone(true, true)[0];

                    for (var j = 0, jLen = this.s.iTableColumns ; j < jLen ; j++) {
                        if ($.inArray(j, aiColumns) === -1) {
                            continue;
                        }

                        var iCloned = $.inArray(aoOriginal[i][j].cell, aCloned);
                        if (iCloned === -1) {
                            var nClone = $(aoOriginal[i][j].cell).clone(true, true)[0];
                            aClones.push(nClone);
                            aCloned.push(aoOriginal[i][j].cell);

                            aRow.push({
                                "cell": nClone,
                                "unique": aoOriginal[i][j].unique
                            });
                        }
                        else {
                            aRow.push({
                                "cell": aClones[iCloned],
                                "unique": aoOriginal[i][j].unique
                            });
                        }
                    }

                    aReturn.push(aRow);
                }

                return aReturn;
            },


            /**
             * Clone the DataTable nodes and place them in the DOM (sized correctly)
             *  @returns {void}
             *  @param   {Object} oClone Object containing the header, footer and body cloned DOM elements
             *  @param   {Object} oGrid Grid object containing the display grid elements for the cloned
             *                    column (left or right)
             *  @param   {Array} aiColumns Column indexes which should be operated on from the DataTable
             *  @param   {Boolean} bAll Indicate if the header and footer should be updated as well (true)
             *  @private
             */
            "_fnClone": function (oClone, oGrid, aiColumns, bAll) {
                var that = this,
                    i, iLen, j, jLen, jq, nTarget, iColumn, nClone, iIndex, aoCloneLayout,
                    jqCloneThead, aoFixedHeader,
                    dt = this.s.dt;

                /*
                 * Header
                 */
                if (bAll) {
                    if (oClone.header !== null) {
                        oClone.header.parentNode.removeChild(oClone.header);
                    }
                    oClone.header = $(this.dom.header).clone(true, true)[0];
                    oClone.header.className += " DTFC_Cloned";
                    oClone.header.style.width = "100%";
                    oGrid.head.appendChild(oClone.header);

                    /* Copy the DataTables layout cache for the header for our floating column */
                    aoCloneLayout = this._fnCopyLayout(dt.aoHeader, aiColumns);
                    jqCloneThead = $('>thead', oClone.header);
                    jqCloneThead.empty();

                    /* Add the created cloned TR elements to the table */
                    for (i = 0, iLen = aoCloneLayout.length ; i < iLen ; i++) {
                        jqCloneThead[0].appendChild(aoCloneLayout[i].nTr);
                    }

                    /* Use the handy _fnDrawHead function in DataTables to do the rowspan/colspan
                     * calculations for us
                     */
                    dt.oApi._fnDrawHead(dt, aoCloneLayout, true);
                }
                else {
                    /* To ensure that we copy cell classes exactly, regardless of colspan, multiple rows
                     * etc, we make a copy of the header from the DataTable again, but don't insert the
                     * cloned cells, just copy the classes across. To get the matching layout for the
                     * fixed component, we use the DataTables _fnDetectHeader method, allowing 1:1 mapping
                     */
                    aoCloneLayout = this._fnCopyLayout(dt.aoHeader, aiColumns);
                    aoFixedHeader = [];

                    dt.oApi._fnDetectHeader(aoFixedHeader, $('>thead', oClone.header)[0]);

                    for (i = 0, iLen = aoCloneLayout.length ; i < iLen ; i++) {
                        for (j = 0, jLen = aoCloneLayout[i].length ; j < jLen ; j++) {
                            aoFixedHeader[i][j].cell.className = aoCloneLayout[i][j].cell.className;

                            // If jQuery UI theming is used we need to copy those elements as well
                            $('span.DataTables_sort_icon', aoFixedHeader[i][j].cell).each(function () {
                                this.className = $('span.DataTables_sort_icon', aoCloneLayout[i][j].cell)[0].className;
                            });
                        }
                    }
                }
                this._fnEqualiseHeights('thead', this.dom.header, oClone.header);

                /*
                 * Body
                 */
                if (this.s.sHeightMatch == 'auto') {
                    /* Remove any heights which have been applied already and let the browser figure it out */
                    $('>tbody>tr', that.dom.body).css('height', 'auto');
                }

                if (oClone.body !== null) {
                    oClone.body.parentNode.removeChild(oClone.body);
                    oClone.body = null;
                }

                oClone.body = $(this.dom.body).clone(true)[0];
                oClone.body.className += " DTFC_Cloned";
                oClone.body.style.paddingBottom = dt.oScroll.iBarWidth + "px";
                oClone.body.style.marginBottom = (dt.oScroll.iBarWidth * 2) + "px"; /* For IE */
                if (oClone.body.getAttribute('id') !== null) {
                    oClone.body.removeAttribute('id');
                }

                $('>thead>tr', oClone.body).empty();
                $('>tfoot', oClone.body).remove();

                var nBody = $('tbody', oClone.body)[0];
                $(nBody).empty();
                if (dt.aiDisplay.length > 0) {
                    /* Copy the DataTables' header elements to force the column width in exactly the
                     * same way that DataTables does it - have the header element, apply the width and
                     * colapse it down
                     */
                    var nInnerThead = $('>thead>tr', oClone.body)[0];
                    for (iIndex = 0 ; iIndex < aiColumns.length ; iIndex++) {
                        iColumn = aiColumns[iIndex];

                        nClone = $(dt.aoColumns[iColumn].nTh).clone(true)[0];
                        nClone.innerHTML = "";

                        var oStyle = nClone.style;
                        oStyle.paddingTop = "0";
                        oStyle.paddingBottom = "0";
                        oStyle.borderTopWidth = "0";
                        oStyle.borderBottomWidth = "0";
                        oStyle.height = 0;
                        oStyle.width = that.s.aiInnerWidths[iColumn] + "px";

                        nInnerThead.appendChild(nClone);
                    }

                    /* Add in the tbody elements, cloning form the master table */
                    $('>tbody>tr', that.dom.body).each(function (z) {
                        var n = this.cloneNode(false);
                        n.removeAttribute('id');
                        var i = that.s.dt.oFeatures.bServerSide === false ?
                            that.s.dt.aiDisplay[that.s.dt._iDisplayStart + z] : z;
                        var aTds = that.s.dt.aoData[i].anCells || $(this).children('td, th');

                        for (iIndex = 0 ; iIndex < aiColumns.length ; iIndex++) {
                            iColumn = aiColumns[iIndex];

                            if (aTds.length > 0) {
                                nClone = $(aTds[iColumn]).clone(true, true)[0];
                                n.appendChild(nClone);
                            }
                        }
                        nBody.appendChild(n);
                    });
                }
                else {
                    $('>tbody>tr', that.dom.body).each(function (z) {
                        nClone = this.cloneNode(true);
                        nClone.className += ' DTFC_NoData';
                        $('td', nClone).html('');
                        nBody.appendChild(nClone);
                    });
                }

                oClone.body.style.width = "100%";
                oClone.body.style.margin = "0";
                oClone.body.style.padding = "0";

                // Interop with Scroller - need to use a height forcing element in the
                // scrolling area in the same way that Scroller does in the body scroll.
                if (dt.oScroller !== undefined) {
                    var scrollerForcer = dt.oScroller.dom.force;

                    if (!oGrid.forcer) {
                        oGrid.forcer = scrollerForcer.cloneNode(true);
                        oGrid.liner.appendChild(oGrid.forcer);
                    }
                    else {
                        oGrid.forcer.style.height = scrollerForcer.style.height;
                    }
                }

                oGrid.liner.appendChild(oClone.body);

                this._fnEqualiseHeights('tbody', that.dom.body, oClone.body);

                /*
                 * Footer
                 */
                if (dt.nTFoot !== null) {
                    if (bAll) {
                        if (oClone.footer !== null) {
                            oClone.footer.parentNode.removeChild(oClone.footer);
                        }
                        oClone.footer = $(this.dom.footer).clone(true, true)[0];
                        oClone.footer.className += " DTFC_Cloned";
                        oClone.footer.style.width = "100%";
                        oGrid.foot.appendChild(oClone.footer);

                        /* Copy the footer just like we do for the header */
                        aoCloneLayout = this._fnCopyLayout(dt.aoFooter, aiColumns);
                        var jqCloneTfoot = $('>tfoot', oClone.footer);
                        jqCloneTfoot.empty();

                        for (i = 0, iLen = aoCloneLayout.length ; i < iLen ; i++) {
                            jqCloneTfoot[0].appendChild(aoCloneLayout[i].nTr);
                        }
                        dt.oApi._fnDrawHead(dt, aoCloneLayout, true);
                    }
                    else {
                        aoCloneLayout = this._fnCopyLayout(dt.aoFooter, aiColumns);
                        var aoCurrFooter = [];

                        dt.oApi._fnDetectHeader(aoCurrFooter, $('>tfoot', oClone.footer)[0]);

                        for (i = 0, iLen = aoCloneLayout.length ; i < iLen ; i++) {
                            for (j = 0, jLen = aoCloneLayout[i].length ; j < jLen ; j++) {
                                aoCurrFooter[i][j].cell.className = aoCloneLayout[i][j].cell.className;
                            }
                        }
                    }
                    this._fnEqualiseHeights('tfoot', this.dom.footer, oClone.footer);
                }

                /* Equalise the column widths between the header footer and body - body get's priority */
                var anUnique = dt.oApi._fnGetUniqueThs(dt, $('>thead', oClone.header)[0]);
                $(anUnique).each(function (i) {
                    iColumn = aiColumns[i];
                    this.style.width = that.s.aiInnerWidths[iColumn] + "px";
                });

                if (that.s.dt.nTFoot !== null) {
                    anUnique = dt.oApi._fnGetUniqueThs(dt, $('>tfoot', oClone.footer)[0]);
                    $(anUnique).each(function (i) {
                        iColumn = aiColumns[i];
                        this.style.width = that.s.aiInnerWidths[iColumn] + "px";
                    });
                }
            },


            /**
             * From a given table node (THEAD etc), get a list of TR direct child elements
             *  @param   {Node} nIn Table element to search for TR elements (THEAD, TBODY or TFOOT element)
             *  @returns {Array} List of TR elements found
             *  @private
             */
            "_fnGetTrNodes": function (nIn) {
                var aOut = [];
                for (var i = 0, iLen = nIn.childNodes.length ; i < iLen ; i++) {
                    if (nIn.childNodes[i].nodeName.toUpperCase() == "TR") {
                        aOut.push(nIn.childNodes[i]);
                    }
                }
                return aOut;
            },


            /**
             * Equalise the heights of the rows in a given table node in a cross browser way
             *  @returns {void}
             *  @param   {String} nodeName Node type - thead, tbody or tfoot
             *  @param   {Node} original Original node to take the heights from
             *  @param   {Node} clone Copy the heights to
             *  @private
             */
            "_fnEqualiseHeights": function (nodeName, original, clone) {
                if (this.s.sHeightMatch == 'none' && nodeName !== 'thead' && nodeName !== 'tfoot') {
                    return;
                }

                var that = this,
                    i, iLen, iHeight, iHeight2, iHeightOriginal, iHeightClone,
                    rootOriginal = original.getElementsByTagName(nodeName)[0],
                    rootClone = clone.getElementsByTagName(nodeName)[0],
                    jqBoxHack = $('>' + nodeName + '>tr:eq(0)', original).children(':first'),
                    iBoxHack = jqBoxHack.outerHeight() - jqBoxHack.height(),
                    anOriginal = this._fnGetTrNodes(rootOriginal),
                    anClone = this._fnGetTrNodes(rootClone),
                    heights = [];

                for (i = 0, iLen = anClone.length ; i < iLen ; i++) {
                    iHeightOriginal = anOriginal[i].offsetHeight;
                    iHeightClone = anClone[i].offsetHeight;
                    iHeight = iHeightClone > iHeightOriginal ? iHeightClone : iHeightOriginal;

                    if (this.s.sHeightMatch == 'semiauto') {
                        anOriginal[i]._DTTC_iHeight = iHeight;
                    }

                    heights.push(iHeight);
                }

                for (i = 0, iLen = anClone.length ; i < iLen ; i++) {
                    anClone[i].style.height = heights[i] + "px";
                    anOriginal[i].style.height = heights[i] + "px";
                }
            }
        };



        /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
         * Statics
         * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

        /**
         * FixedColumns default settings for initialisation
         *  @name FixedColumns.defaults
         *  @namespace
         *  @static
         */
        FixedColumns.defaults = /** @lends FixedColumns.defaults */{
            /**
             * Number of left hand columns to fix in position
             *  @type     int
             *  @default  1
             *  @static
             *  @example
             *      var  = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      new $.fn.dataTable.fixedColumns( table, {
             *          "leftColumns": 2
             *      } );
             */
            "iLeftColumns": 1,

            /**
             * Number of right hand columns to fix in position
             *  @type     int
             *  @default  0
             *  @static
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      new $.fn.dataTable.fixedColumns( table, {
             *          "rightColumns": 1
             *      } );
             */
            "iRightColumns": 0,

            /**
             * Draw callback function which is called when FixedColumns has redrawn the fixed assets
             *  @type     function(object, object):void
             *  @default  null
             *  @static
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      new $.fn.dataTable.fixedColumns( table, {
             *          "drawCallback": function () {
             *	            alert( "FixedColumns redraw" );
             *	        }
             *      } );
             */
            "fnDrawCallback": null,

            /**
             * Height matching algorthim to use. This can be "none" which will result in no height
             * matching being applied by FixedColumns (height matching could be forced by CSS in this
             * case), "semiauto" whereby the height calculation will be performed once, and the result
             * cached to be used again (fnRecalculateHeight can be used to force recalculation), or
             * "auto" when height matching is performed on every draw (slowest but must accurate)
             *  @type     string
             *  @default  semiauto
             *  @static
             *  @example
             *      var table = $('#example').dataTable( {
             *          "scrollX": "100%"
             *      } );
             *      new $.fn.dataTable.fixedColumns( table, {
             *          "heightMatch": "auto"
             *      } );
             */
            "sHeightMatch": "semiauto"
        };




        /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
         * Constants
         * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

        /**
         * FixedColumns version
         *  @name      FixedColumns.version
         *  @type      String
         *  @default   See code
         *  @static
         */
        FixedColumns.version = "3.0.4";



        /* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
         * Fired events (for documentation)
         * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */


        /**
         * Event fired whenever FixedColumns redraws the fixed columns (i.e. clones the table elements from the main DataTable). This will occur whenever the DataTable that the FixedColumns instance is attached does its own draw.
         * @name FixedColumns#draw.dtfc
         * @event
         * @param {event} e jQuery event object
         * @param {object} o Event parameters from FixedColumns
         * @param {object} o.leftClone Instance's object dom.clone.left for easy reference. This object contains references to the left fixed clumn column's nodes
         * @param {object} o.rightClone Instance's object dom.clone.right for easy reference. This object contains references to the right fixed clumn column's nodes
         */


        // Make FixedColumns accessible from the DataTables instance
        $.fn.dataTable.FixedColumns = FixedColumns;
        $.fn.DataTable.FixedColumns = FixedColumns;


        return FixedColumns;
    }; // /factory


    // Define as an AMD module if possible
    if (typeof define === 'function' && define.amd) {
        define(['jquery.ba-resize', 'datatables'], factory);
    }
    else if (typeof exports === 'object') {
        // Node/CommonJS
        factory(require('jquery'), require('datatables'));
    }
    else if (jQuery && !jQuery.fn.dataTable.FixedColumns) {
        // Otherwise simply initialise as normal, stopping multiple evaluation
        factory(jQuery, jQuery.fn.dataTable);
    }


})(window, document);

﻿//Date.prototype.isLeapYear 判断闰年
//Date.prototype.Format 日期格式化
//Date.prototype.DateAdd 日期计算
//Date.prototype.DateDiff 比较日期差
//Date.prototype.toString 日期转字符串
//Date.prototype.toArray 日期分割为数组
//Date.prototype.DatePart 取日期的部分信息
//Date.prototype.MaxDayOfDate 取日期所在月的最大天数
//Date.prototype.WeekNumOfYear 判断日期所在年的第几周
//StringToDate 字符串转日期型
//IsValidDate 验证日期有效性
//CheckDateTime 完整日期时间检查
//daysBetween 日期天数差

//少量调用案例
//var temp = $('#txtDate').val()
//var nowDate = new Date(temp);
//nowDate = new Date(nowDate.getFullYear() + "/" + (nowDate.getMonth() + 1) + "/" + "01")
////alert(Date.parse(nowDate));

//nowDate = nowDate.DateAdd('m', 1).DateAdd('d', -1);
//alert(nowDate.Format("YYYY/MM/DD"));
//var strDate = nowDate.getFullYear() + "/" + (nowDate.getMonth() + 1) + "/" + nowDate.getDate();
//alert(strDate);

////---------------------------------------------------  
//// 判断闰年  
////---------------------------------------------------  
//Date.prototype.isLeapYear = function()   
//{   
//    return (0==this.getYear()%4&&((this.getYear()%100!=0)||(this.getYear()%400==0)));   
//}   
  
function stringToDate(DateStr) {
    if (DateStr.indexOf(".") > 0) {
        DateStr = DateStr.substr(0, DateStr.indexOf("."));
    }
    DateStr = DateStr.replace(/-/g, '/');

    var converted = Date.parse(DateStr);

    //alert(DateStr.substr(0, 4) + "/" + DateStr.substr(5, 2) + "/" + DateStr.substr(8, 2));
    var myDate = new Date(converted);
    if (isNaN(converted)) {

        myDate = new Date();
    }
    return myDate;
}

//---------------------------------------------------  
// 日期格式化  
// 格式 YYYY/yyyy/YY/yy 表示年份  
// MM/M 月份  
// W/w 星期  
// dd/DD/d/D 日期  
// hh/HH/h/H 时间  
// mm/m 分钟  
// ss/SS/s/S 秒  
//---------------------------------------------------  
Date.prototype.format = function(formatStr)   
{   
    var str = formatStr;   
    var Week = ['日','一','二','三','四','五','六'];  
  
    str=str.replace(/yyyy|YYYY/,this.getFullYear());   
    str=str.replace(/yy|YY/,(this.getYear() % 100)>9?(this.getYear() % 100).toString():'0' + (this.getYear() % 100));   
  
    str = str.replace(/MM/, (this.getMonth() + 1) > 9 ? (this.getMonth() + 1).toString() : '0' + (this.getMonth() + 1));
    str = str.replace(/M/g, (this.getMonth() + 1));
  
    str=str.replace(/w|W/g,Week[this.getDay()]);   
  
    str=str.replace(/dd|DD/,this.getDate()>9?this.getDate().toString():'0' + this.getDate());   
    str=str.replace(/d|D/g,this.getDate());   
  
    str=str.replace(/hh|HH/,this.getHours()>9?this.getHours().toString():'0' + this.getHours());   
    str=str.replace(/h|H/g,this.getHours());   
    str=str.replace(/mm/,this.getMinutes()>9?this.getMinutes().toString():'0' + this.getMinutes());   
    str=str.replace(/m/g,this.getMinutes());   
  
    str=str.replace(/ss|SS/,this.getSeconds()>9?this.getSeconds().toString():'0' + this.getSeconds());   
    str=str.replace(/s|S/g,this.getSeconds());   
  
    return str;   
}   
  
////+---------------------------------------------------  
////| 求两个时间的天数差 日期格式为 YYYY-MM-dd   
////+---------------------------------------------------  
//function daysBetween(DateOne,DateTwo)  
//{   
//    var OneMonth = DateOne.substring(5,DateOne.lastIndexOf ('-'));  
//    var OneDay = DateOne.substring(DateOne.length,DateOne.lastIndexOf ('-')+1);  
//    var OneYear = DateOne.substring(0,DateOne.indexOf ('-'));  
  
//    var TwoMonth = DateTwo.substring(5,DateTwo.lastIndexOf ('-'));  
//    var TwoDay = DateTwo.substring(DateTwo.length,DateTwo.lastIndexOf ('-')+1);  
//    var TwoYear = DateTwo.substring(0,DateTwo.indexOf ('-'));  
  
//    var cha=((Date.parse(OneMonth+'/'+OneDay+'/'+OneYear)- Date.parse(TwoMonth+'/'+TwoDay+'/'+TwoYear))/86400000);   
//    return Math.abs(cha);  
//}  
  
  
//+---------------------------------------------------  
//| 日期计算  
//+---------------------------------------------------  
Date.prototype.dateAdd = function(strInterval, Number) {   
    var dtTmp = this;  
    switch (strInterval) {   
        case 's' :return new Date(Date.parse(dtTmp) + (1000 * Number));  
        case 'n' :return new Date(Date.parse(dtTmp) + (60000 * Number));  
        case 'h' :return new Date(Date.parse(dtTmp) + (3600000 * Number));  
        case 'd' :return new Date(Date.parse(dtTmp) + (86400000 * Number));  
        case 'w' :return new Date(Date.parse(dtTmp) + ((86400000 * 7) * Number));  
        case 'q' :return new Date(dtTmp.getFullYear(), (dtTmp.getMonth()) + Number*3, dtTmp.getDate(), dtTmp.getHours(), dtTmp.getMinutes(), dtTmp.getSeconds());  
        case 'm' :return new Date(dtTmp.getFullYear(), (dtTmp.getMonth()) + Number, dtTmp.getDate(), dtTmp.getHours(), dtTmp.getMinutes(), dtTmp.getSeconds());  
        case 'y' :return new Date((dtTmp.getFullYear() + Number), dtTmp.getMonth(), dtTmp.getDate(), dtTmp.getHours(), dtTmp.getMinutes(), dtTmp.getSeconds());  
    }  
}  
  
//+---------------------------------------------------  
//| 比较日期差 dtEnd 格式为日期型或者 有效日期格式字符串  
//+---------------------------------------------------  
Date.prototype.DateDiff = function(strInterval, dtEnd) {   
    var dtStart = this;  
    if (typeof dtEnd == 'string' )//如果是字符串转换为日期型  
    {   
        dtEnd = StringToDate(dtEnd);  
    }  
    switch (strInterval) {   
        case 's' :return parseInt((dtEnd - dtStart) / 1000);  
        case 'n' :return parseInt((dtEnd - dtStart) / 60000);  
        case 'h' :return parseInt((dtEnd - dtStart) / 3600000);  
        case 'd' :return parseInt((dtEnd - dtStart) / 86400000);  
        case 'w' :return parseInt((dtEnd - dtStart) / (86400000 * 7));  
        case 'm' :return (dtEnd.getMonth()+1)+((dtEnd.getFullYear()-dtStart.getFullYear())*12) - (dtStart.getMonth()+1);  
        case 'y' :return dtEnd.getFullYear() - dtStart.getFullYear();  
    }  
}  
  
////+---------------------------------------------------  
////| 日期输出字符串，重载了系统的toString方法  
////+---------------------------------------------------  
//Date.prototype.toString = function(showWeek)  
//{   
//    var myDate= this;  
//    var str = myDate.toLocaleDateString();  
//    if (showWeek)  
//    {   
//        var Week = ['日','一','二','三','四','五','六'];  
//        str += ' 星期' + Week[myDate.getDay()];  
//    }  
//    return str;  
//}  
  
////+---------------------------------------------------  
////| 日期合法性验证  
////| 格式为：YYYY-MM-DD或YYYY/MM/DD  
////+---------------------------------------------------  
//function IsValidDate(DateStr)   
//{   
//    var sDate=DateStr.replace(/(^\s+|\s+$)/g,''); //去两边空格;   
//    if(sDate=='') return true;   
//    //如果格式满足YYYY-(/)MM-(/)DD或YYYY-(/)M-(/)DD或YYYY-(/)M-(/)D或YYYY-(/)MM-(/)D就替换为''   
//    //数据库中，合法日期可以是:YYYY-MM/DD(2003-3/21),数据库会自动转换为YYYY-MM-DD格式   
//    var s = sDate.replace(/[\d]{ 4,4 }[\-/]{ 1 }[\d]{ 1,2 }[\-/]{ 1 }[\d]{ 1,2 }/g,'');   
//    if (s=='') //说明格式满足YYYY-MM-DD或YYYY-M-DD或YYYY-M-D或YYYY-MM-D   
//    {   
//        var t=new Date(sDate.replace(/\-/g,'/'));   
//        var ar = sDate.split(/[-/:]/);   
//        if(ar[0] != t.getYear() || ar[1] != t.getMonth()+1 || ar[2] != t.getDate())   
//        {   
//            //alert('错误的日期格式！格式为：YYYY-MM-DD或YYYY/MM/DD。注意闰年。');   
//            return false;   
//        }   
//    }   
//    else   
//    {   
//        //alert('错误的日期格式！格式为：YYYY-MM-DD或YYYY/MM/DD。注意闰年。');   
//        return false;   
//    }   
//    return true;   
//}   
  
////+---------------------------------------------------  
////| 日期时间检查  
////| 格式为：YYYY-MM-DD HH:MM:SS  
////+---------------------------------------------------  
//function CheckDateTime(str)  
//{   
//    var reg = /^(\d+)-(\d{ 1,2 })-(\d{ 1,2 }) (\d{ 1,2 }):(\d{ 1,2 }):(\d{ 1,2 })$/;   
//    var r = str.match(reg);   
//    if(r==null)return false;   
//    r[2]=r[2]-1;   
//    var d= new Date(r[1],r[2],r[3],r[4],r[5],r[6]);   
//    if(d.getFullYear()!=r[1])return false;   
//    if(d.getMonth()!=r[2])return false;   
//    if(d.getDate()!=r[3])return false;   
//    if(d.getHours()!=r[4])return false;   
//    if(d.getMinutes()!=r[5])return false;   
//    if(d.getSeconds()!=r[6])return false;   
//    return true;   
//}   
  
////+---------------------------------------------------  
////| 把日期分割成数组  
////+---------------------------------------------------  
//Date.prototype.toArray = function()  
//{   
//    var myDate = this;  
//    var myArray = Array();  
//    myArray[0] = myDate.getFullYear();  
//    myArray[1] = myDate.getMonth();  
//    myArray[2] = myDate.getDate();  
//    myArray[3] = myDate.getHours();  
//    myArray[4] = myDate.getMinutes();  
//    myArray[5] = myDate.getSeconds();  
//    return myArray;  
//}  
  
////+---------------------------------------------------  
////| 取得日期数据信息  
////| 参数 interval 表示数据类型  
////| y 年 m月 d日 w星期 ww周 h时 n分 s秒  
////+---------------------------------------------------  
//Date.prototype.DatePart = function(interval)  
//{   
//    var myDate = this;  
//    var partStr='';  
//    var Week = ['日','一','二','三','四','五','六'];  
//    switch (interval)  
//    {   
//        case 'y' :partStr = myDate.getFullYear();break;  
//        case 'm' :partStr = myDate.getMonth()+1;break;  
//        case 'd' :partStr = myDate.getDate();break;  
//        case 'w' :partStr = Week[myDate.getDay()];break;  
//        case 'ww' :partStr = myDate.WeekNumOfYear();break;  
//        case 'h' :partStr = myDate.getHours();break;  
//        case 'n' :partStr = myDate.getMinutes();break;  
//        case 's' :partStr = myDate.getSeconds();break;  
//    }  
//    return partStr;  
//}  
  
////+---------------------------------------------------  
////| 取得当前日期所在月的最大天数  
////+---------------------------------------------------  
//Date.prototype.MaxDayOfDate = function()  
//{   
//    var myDate = this;  
//    var ary = myDate.toArray();  
//    var date1 = (new Date(ary[0],ary[1]+1,1));  
//    var date2 = date1.dateAdd(1,'m',1);  
//    var result = dateDiff(date1.Format('yyyy-MM-dd'),date2.Format('yyyy-MM-dd'));  
//    return result;  
//}  
  
//////+---------------------------------------------------  
//////| 取得当前日期所在周是一年中的第几周  
//////+---------------------------------------------------  
////Date.prototype.WeekNumOfYear = function()  
////{   
////    var myDate = this;  
////    var ary = myDate.toArray();  
////    var year = ary[0];  
////    var month = ary[1]+1;  
////    var day = ary[2];  
////    document.write('< script language=VBScript\> \n');  
////    document.write('myDate = DateValue(''+month+'-'+day+'-'+year+'') \n');  
////    document.write('result = DatePart('ww', myDate) \n');  
////    document.write(' \n');  
////    return result;  
////}  
  
////+---------------------------------------------------  
////| 字符串转成日期类型   
////| 格式 MM/dd/YYYY MM-dd-YYYY YYYY/MM/dd YYYY-MM-dd  
////+---------------------------------------------------  
//function StringToDate(DateStr)  
//{   
  
//    var converted = Date.parse(DateStr);  
//    var myDate = new Date(converted);  
//    if (isNaN(myDate))  
//    {   
//        //var delimCahar = DateStr.indexOf('/')!=-1?'/':'-';  
//        var arys= DateStr.split('-');  
//        myDate = new Date(arys[0],--arys[1],arys[2]);  
//    }  
//    return myDate;  
//}  
﻿/**
 * jquery.numberformatter - Formatting/Parsing Numbers in jQuery
 * 
 * Written by
 * Michael Abernethy (mike@abernethysoft.com),
 * Andrew Parry (aparry0@gmail.com)
 *
 * Dual licensed under the MIT (MIT-LICENSE.txt)
 * and GPL (GPL-LICENSE.txt) licenses.
 *
 * @author Michael Abernethy, Andrew Parry
 * @version 1.2.4-RELEASE ($Id$)
 * 
 * Dependencies
 * 
 * jQuery (http://jquery.com)
 * jshashtable (http://www.timdown.co.uk/jshashtable)
 * 
 * Notes & Thanks
 * 
 * many thanks to advweb.nanasi.jp for his bug fixes
 * jsHashtable is now used also, so thanks to the author for that excellent little class.
 *
 * This plugin can be used to format numbers as text and parse text as Numbers
 * Because we live in an international world, we cannot assume that everyone
 * uses "," to divide thousands, and "." as a decimal point.
 *
 * As of 1.2 the way this plugin works has changed slightly, parsing text to a number
 * has 1 set of functions, formatting a number to text has it's own. Before things
 * were a little confusing, so I wanted to separate the 2 out more.
 *
 *
 * jQuery extension functions:
 *
 * formatNumber(options, writeBack, giveReturnValue) - Reads the value from the subject, parses to
 * a Javascript Number object, then formats back to text using the passed options and write back to
 * the subject.
 * 
 * parseNumber(options) - Parses the value in the subject to a Number object using the passed options
 * to decipher the actual number from the text, then writes the value as text back to the subject.
 * 
 * 
 * Generic functions:
 * 
 * formatNumber(numberString, options) - Takes a plain number as a string (e.g. '1002.0123') and returns
 * a string of the given format options.
 * 
 * parseNumber(numberString, options) - Takes a number as text that is formatted the same as the given
 * options then and returns it as a plain Number object.
 * 
 * To achieve the old way of combining parsing and formatting to keep say a input field always formatted
 * to a given format after it has lost focus you'd simply use a combination of the functions.
 * 
 * e.g.
 * $("#salary").blur(function(){
 *      $(this).parseNumber({format:"#,###.00", locale:"us"});
 *      $(this).formatNumber({format:"#,###.00", locale:"us"});
 * });
 *
 * The syntax for the formatting is:
 * 0 = Digit
 * # = Digit, zero shows as absent
 * . = Decimal separator
 * - = Negative sign
 * , = Grouping Separator
 * % = Percent (multiplies the number by 100)
 * 
 * For example, a format of "#,###.00" and text of 4500.20 will
 * display as "4.500,20" with a locale of "de", and "4,500.20" with a locale of "us"
 *
 *
 * As of now, the only acceptable locales are 
 * Arab Emirates -> "ae"
 * Australia -> "au"
 * Austria -> "at"
 * Brazil -> "br"
 * Canada -> "ca"
 * China -> "cn"
 * Czech -> "cz"
 * Denmark -> "dk"
 * Egypt -> "eg"
 * Finland -> "fi"
 * France  -> "fr"
 * Germany -> "de"
 * Greece -> "gr"
 * Great Britain -> "gb"
 * Hong Kong -> "hk"
 * India -> "in"
 * Israel -> "il"
 * Japan -> "jp"
 * Russia -> "ru"
 * South Korea -> "kr"
 * Spain -> "es"
 * Sweden -> "se"
 * Switzerland -> "ch"
 * Taiwan -> "tw"
 * Thailand -> "th"
 * United States -> "us"
 * Vietnam -> "vn"
 **/
var Hashtable = function (t) { function n(t) { return typeof t == p ? t : "" + t } function e(t) { var r; return typeof t == p ? t : typeof t.hashCode == y ? (r = t.hashCode(), typeof r == p ? r : e(r)) : n(t) } function r(t, n) { for (var e in n) n.hasOwnProperty(e) && (t[e] = n[e]) } function i(t, n) { return t.equals(n) } function u(t, n) { return typeof n.equals == y ? n.equals(t) : t === n } function o(n) { return function (e) { if (null === e) throw new Error("null is not a valid " + n); if (e === t) throw new Error(n + " must not be undefined") } } function s(t, n, e, r) { this[0] = t, this.entries = [], this.addEntry(n, e), null !== r && (this.getEqualityFunction = function () { return r }) } function a(t) { return function (n) { for (var e, r = this.entries.length, i = this.getEqualityFunction(n) ; r--;) if (e = this.entries[r], i(n, e[0])) switch (t) { case E: return !0; case K: return e; case q: return [r, e[1]] } return !1 } } function l(t) { return function (n) { for (var e = n.length, r = 0, i = this.entries, u = i.length; u > r; ++r) n[e + r] = i[r][t] } } function f(t, n) { for (var e, r = t.length; r--;) if (e = t[r], n === e[0]) return r; return null } function h(t, n) { var e = t[n]; return e && e instanceof s ? e : null } function c() { var n = [], i = {}, u = { replaceDuplicateKey: !0, hashCode: e, equals: null }, o = arguments[0], a = arguments[1]; a !== t ? (u.hashCode = o, u.equals = a) : o !== t && r(u, o); var l = u.hashCode, c = u.equals; this.properties = u, this.put = function (t, e) { g(t), d(e); var r, o, a = l(t), f = null; return r = h(i, a), r ? (o = r.getEntryForKey(t), o ? (u.replaceDuplicateKey && (o[0] = t), f = o[1], o[1] = e) : r.addEntry(t, e)) : (r = new s(a, t, e, c), n.push(r), i[a] = r), f }, this.get = function (t) { g(t); var n = l(t), e = h(i, n); if (e) { var r = e.getEntryForKey(t); if (r) return r[1] } return null }, this.containsKey = function (t) { g(t); var n = l(t), e = h(i, n); return e ? e.containsKey(t) : !1 }, this.containsValue = function (t) { d(t); for (var e = n.length; e--;) if (n[e].containsValue(t)) return !0; return !1 }, this.clear = function () { n.length = 0, i = {} }, this.isEmpty = function () { return !n.length }; var y = function (t) { return function () { for (var e = [], r = n.length; r--;) n[r][t](e); return e } }; this.keys = y("keys"), this.values = y("values"), this.entries = y("getEntries"), this.remove = function (t) { g(t); var e, r = l(t), u = null, o = h(i, r); return o && (u = o.removeEntryForKey(t), null !== u && 0 == o.entries.length && (e = f(n, r), n.splice(e, 1), delete i[r])), u }, this.size = function () { for (var t = 0, e = n.length; e--;) t += n[e].entries.length; return t } } var y = "function", p = "string", v = "undefined"; if (typeof encodeURIComponent == v || Array.prototype.splice === t || Object.prototype.hasOwnProperty === t) return null; var g = o("key"), d = o("value"), E = 0, K = 1, q = 2; return s.prototype = { getEqualityFunction: function (t) { return typeof t.equals == y ? i : u }, getEntryForKey: a(K), getEntryAndIndexForKey: a(q), removeEntryForKey: function (t) { var n = this.getEntryAndIndexForKey(t); return n ? (this.entries.splice(n[0], 1), n[1]) : null }, addEntry: function (t, n) { this.entries.push([t, n]) }, keys: l(0), values: l(1), getEntries: function (t) { for (var n = t.length, e = 0, r = this.entries, i = r.length; i > e; ++e) t[n + e] = r[e].slice(0) }, containsKey: a(E), containsValue: function (t) { for (var n = this.entries, e = n.length; e--;) if (t === n[e][1]) return !0; return !1 } }, c.prototype = { each: function (t) { for (var n, e = this.entries(), r = e.length; r--;) n = e[r], t(n[0], n[1]) }, equals: function (t) { var n, e, r, i = this.size(); if (i == t.size()) { for (n = this.keys() ; i--;) if (e = n[i], r = t.get(e), null === r || r !== this.get(e)) return !1; return !0 } return !1 }, putAll: function (t, n) { for (var e, r, i, u, o = t.entries(), s = o.length, a = typeof n == y; s--;) e = o[s], r = e[0], i = e[1], a && (u = this.get(r)) && (i = n(r, u, i)), this.put(r, i) }, clone: function () { var t = new c(this.properties); return t.putAll(this), t } }, c.prototype.toQueryString = function () { for (var t, e = this.entries(), r = e.length, i = []; r--;) t = e[r], i[r] = encodeURIComponent(n(t[0])) + "=" + encodeURIComponent(n(t[1])); return i.join("&") }, c }();

(function (jQuery) {

    var nfLocales = new Hashtable();

    var nfLocalesLikeUS = ['ae', 'au', 'ca', 'cn', 'eg', 'gb', 'hk', 'il', 'in', 'jp', 'sk', 'th', 'tw', 'us'];
    var nfLocalesLikeDE = ['at', 'br', 'de', 'dk', 'es', 'gr', 'it', 'nl', 'pt', 'tr', 'vn'];
    var nfLocalesLikeFR = ['bg', 'cz', 'fi', 'fr', 'no', 'pl', 'ru', 'se'];
    var nfLocalesLikeCH = ['ch'];

    var nfLocaleFormatting = [[".", ","], [",", "."], [",", " "], [".", "'"]];
    var nfAllLocales = [nfLocalesLikeUS, nfLocalesLikeDE, nfLocalesLikeFR, nfLocalesLikeCH]

    function FormatData(dec, group, neg) {
        this.dec = dec;
        this.group = group;
        this.neg = neg;
    };

    function init() {
        // write the arrays into the hashtable
        for (var localeGroupIdx = 0; localeGroupIdx < nfAllLocales.length; localeGroupIdx++) {
            var localeGroup = nfAllLocales[localeGroupIdx];
            for (var i = 0; i < localeGroup.length; i++) {
                nfLocales.put(localeGroup[i], localeGroupIdx);
            }
        }
    };

    function formatCodes(locale, isFullLocale) {
        if (nfLocales.size() == 0)
            init();

        // default values
        var dec = ".";
        var group = ",";
        var neg = "-";

        if (isFullLocale == false) {
            // Extract and convert to lower-case any language code from a real 'locale' formatted string, if not use as-is
            // (To prevent locale format like : "fr_FR", "en_US", "de_DE", "fr_FR", "en-US", "de-DE")
            if (locale.indexOf('_') != -1)
                locale = locale.split('_')[1].toLowerCase();
            else if (locale.indexOf('-') != -1)
                locale = locale.split('-')[1].toLowerCase();
        }

        // hashtable lookup to match locale with codes
        var codesIndex = nfLocales.get(locale);
        if (codesIndex) {
            var codes = nfLocaleFormatting[codesIndex];
            if (codes) {
                dec = codes[0];
                group = codes[1];
            }
        }
        return new FormatData(dec, group, neg);
    };


    /*  Formatting Methods  */


    /**
     * Formats anything containing a number in standard js number notation.
     * 
     * @param {Object}  options         The formatting options to use
     * @param {Boolean} writeBack       (true) If the output value should be written back to the subject
     * @param {Boolean} giveReturnValue (true) If the function should return the output string
     */
    jQuery.fn.formatNumber = function (options, writeBack, giveReturnValue) {

        return this.each(function () {
            // enforce defaults
            if (writeBack == null)
                writeBack = true;
            if (giveReturnValue == null)
                giveReturnValue = true;

            // get text
            var text;
            if (jQuery(this).is(":input"))
                text = new String(jQuery(this).val());
            else
                text = new String(jQuery(this).text());

            // format
            var returnString = jQuery.formatNumber(text, options);

            // set formatted string back, only if a success
            //          if (returnString) {
            if (writeBack) {
                if (jQuery(this).is(":input"))
                    jQuery(this).val(returnString);
                else
                    jQuery(this).text(returnString);
            }
            if (giveReturnValue)
                return returnString;
            //          }
            //          return '';
        });
    };

    /**
     * First parses a string and reformats it with the given options.
     * 
     * @param {Object} numberString
     * @param {Object} options
     */
    jQuery.formatNumber = function (numberString, options) {
        var options = jQuery.extend({}, jQuery.fn.formatNumber.defaults, options);
        var formatData = formatCodes(options.locale.toLowerCase(), options.isFullLocale);

        var dec = formatData.dec;
        var group = formatData.group;
        var neg = formatData.neg;

        var validFormat = "0#-,.";

        // strip all the invalid characters at the beginning and the end
        // of the format, and we'll stick them back on at the end
        // make a special case for the negative sign "-" though, so 
        // we can have formats like -$23.32
        var prefix = "";
        var negativeInFront = false;
        for (var i = 0; i < options.format.length; i++) {
            if (validFormat.indexOf(options.format.charAt(i)) == -1)
                prefix = prefix + options.format.charAt(i);
            else
                if (i == 0 && options.format.charAt(i) == '-') {
                    negativeInFront = true;
                    continue;
                }
                else
                    break;
        }
        var suffix = "";
        for (var i = options.format.length - 1; i >= 0; i--) {
            if (validFormat.indexOf(options.format.charAt(i)) == -1)
                suffix = options.format.charAt(i) + suffix;
            else
                break;
        }

        options.format = options.format.substring(prefix.length);
        options.format = options.format.substring(0, options.format.length - suffix.length);

        // now we need to convert it into a number
        //while (numberString.indexOf(group) > -1) 
        //  numberString = numberString.replace(group, '');
        //var number = new Number(numberString.replace(dec, ".").replace(neg, "-"));
        var number = new Number(numberString);

        return jQuery._formatNumber(number, options, suffix, prefix, negativeInFront);
    };

    /**
     * Formats a Number object into a string, using the given formatting options
     * 
     * @param {Object} numberString
     * @param {Object} options
     */
    jQuery._formatNumber = function (number, options, suffix, prefix, negativeInFront) {
        var options = jQuery.extend({}, jQuery.fn.formatNumber.defaults, options);
        var formatData = formatCodes(options.locale.toLowerCase(), options.isFullLocale);

        var dec = formatData.dec;
        var group = formatData.group;
        var neg = formatData.neg;

        // check overrides
        if (options.overrideGroupSep != null) {
            group = options.overrideGroupSep;
        }
        if (options.overrideDecSep != null) {
            dec = options.overrideDecSep;
        }
        if (options.overrideNegSign != null) {
            neg = options.overrideNegSign;
        }

        // Check NAN handling
        var forcedToZero = false;
        if (isNaN(number)) {
            if (options.nanForceZero == true) {
                number = 0;
                forcedToZero = true;
            } else {
                return '';
            }
        }

        // special case for percentages
        if (options.isPercentage == true || (options.autoDetectPercentage && suffix.charAt(suffix.length - 1) == '%')) {
            number = number * 100;
        }

        var returnString = "";
        if (options.format.indexOf(".") > -1) {
            var decimalPortion = dec;
            var decimalFormat = options.format.substring(options.format.lastIndexOf(".") + 1);

            // round or truncate number as needed
            if (options.round == true)
                number = new Number(number.toFixed(decimalFormat.length));
            else {
                var numStr = number.toString();
                if (numStr.lastIndexOf('.') > 0) {
                    numStr = numStr.substring(0, numStr.lastIndexOf('.') + decimalFormat.length + 1);
                }
                number = new Number(numStr);
            }

            var decimalValue = new Number(number.toString().substring(number.toString().indexOf('.')));
            decimalString = new String(decimalValue.toFixed(decimalFormat.length));
            decimalString = decimalString.substring(decimalString.lastIndexOf('.') + 1);
            for (var i = 0; i < decimalFormat.length; i++) {
                if (decimalFormat.charAt(i) == '#' && decimalString.charAt(i) != '0') {
                    decimalPortion += decimalString.charAt(i);
                    continue;
                } else if (decimalFormat.charAt(i) == '#' && decimalString.charAt(i) == '0') {
                    var notParsed = decimalString.substring(i);
                    if (notParsed.match('[1-9]')) {
                        decimalPortion += decimalString.charAt(i);
                        continue;
                    } else
                        break;
                } else if (decimalFormat.charAt(i) == "0")
                    decimalPortion += decimalString.charAt(i);
            }
            returnString += decimalPortion
        } else
            number = Math.round(number);

        var ones = Math.floor(number);
        if (number < 0)
            ones = Math.ceil(number);

        var onesFormat = "";
        if (options.format.indexOf(".") == -1)
            onesFormat = options.format;
        else
            onesFormat = options.format.substring(0, options.format.indexOf("."));

        var onePortion = "";
        if (!(ones == 0 && onesFormat.substr(onesFormat.length - 1) == '#') || forcedToZero) {
            // find how many digits are in the group
            var oneText = new String(Math.abs(ones));
            var groupLength = 9999;
            if (onesFormat.lastIndexOf(",") != -1)
                groupLength = onesFormat.length - onesFormat.lastIndexOf(",") - 1;
            var groupCount = 0;
            for (var i = oneText.length - 1; i > -1; i--) {
                onePortion = oneText.charAt(i) + onePortion;
                groupCount++;
                if (groupCount == groupLength && i != 0) {
                    onePortion = group + onePortion;
                    groupCount = 0;
                }
            }

            // account for any pre-data padding
            if (onesFormat.length > onePortion.length) {
                var padStart = onesFormat.indexOf('0');
                if (padStart != -1) {
                    var padLen = onesFormat.length - padStart;

                    // pad to left with 0's or group char
                    var pos = onesFormat.length - onePortion.length - 1;
                    while (onePortion.length < padLen) {
                        var padChar = onesFormat.charAt(pos);
                        // replace with real group char if needed
                        if (padChar == ',')
                            padChar = group;
                        onePortion = padChar + onePortion;
                        pos--;
                    }
                }
            }
        }

        if (!onePortion && onesFormat.indexOf('0', onesFormat.length - 1) !== -1)
            onePortion = '0';

        returnString = onePortion + returnString;

        // handle special case where negative is in front of the invalid characters
        if (number < 0 && negativeInFront && prefix.length > 0)
            prefix = neg + prefix;
        else if (number < 0)
            returnString = neg + returnString;

        if (!options.decimalSeparatorAlwaysShown) {
            if (returnString.lastIndexOf(dec) == returnString.length - 1) {
                returnString = returnString.substring(0, returnString.length - 1);
            }
        }
        returnString = prefix + returnString + suffix;
        return returnString;
    };


    /*  Parsing Methods */


    /**
     * Parses a number of given format from the element and returns a Number object.
     * @param {Object} options
     */
    jQuery.fn.parseNumber = function (options, writeBack, giveReturnValue) {
        // enforce defaults
        if (writeBack == null)
            writeBack = true;
        if (giveReturnValue == null)
            giveReturnValue = true;

        // get text
        var text;
        if (jQuery(this).is(":input"))
            text = new String(jQuery(this).val());
        else
            text = new String(jQuery(this).text());

        // parse text
        var number = jQuery.parseNumber(text, options);

        if (number) {
            if (writeBack) {
                if (jQuery(this).is(":input"))
                    jQuery(this).val(number.toString());
                else
                    jQuery(this).text(number.toString());
            }
            if (giveReturnValue)
                return number;
        }
    };

    /**
     * Parses a string of given format into a Number object.
     * 
     * @param {Object} string
     * @param {Object} options
     */
    jQuery.parseNumber = function (numberString, options) {
        var options = jQuery.extend({}, jQuery.fn.parseNumber.defaults, options);
        var formatData = formatCodes(options.locale.toLowerCase(), options.isFullLocale);

        var dec = formatData.dec;
        var group = formatData.group;
        var neg = formatData.neg;

        // check overrides
        if (options.overrideGroupSep != null) {
            group = options.overrideGroupSep;
        }
        if (options.overrideDecSep != null) {
            dec = options.overrideDecSep;
        }
        if (options.overrideNegSign != null) {
            neg = options.overrideNegSign;
        }

        var valid = "1234567890";
        var validOnce = '.-';
        var strictMode = options.strict;

        // now we need to convert it into a number
        while (numberString.indexOf(group) > -1) {
            numberString = numberString.replace(group, '');
        }
        numberString = numberString.replace(dec, '.').replace(neg, '-');
        var validText = '';
        var hasPercent = false;

        if (options.isPercentage == true || (options.autoDetectPercentage && numberString.charAt(numberString.length - 1) == '%')) {
            hasPercent = true;
        }

        for (var i = 0; i < numberString.length; i++) {
            if (valid.indexOf(numberString.charAt(i)) > -1) {
                validText = validText + numberString.charAt(i);
            } else if (validOnce.indexOf(numberString.charAt(i)) > -1) {
                validText = validText + numberString.charAt(i);
                validOnce = validOnce.replace(numberString.charAt(i), '');
            } else {
                if (options.allowPostfix) {
                    // treat anything after this point (inclusive) as a postfix
                    break;
                } else if (strictMode) {
                    // abort and force the text to NaN as it's not strictly valid
                    validText = 'NaN';
                    break;
                }
            }
        }
        var number = new Number(validText);
        if (hasPercent) {
            number = number / 100;
            var decimalPos = validText.indexOf('.');
            if (decimalPos != -1) {
                var decimalPoints = validText.length - decimalPos - 1;
                number = number.toFixed(decimalPoints + 2);
            } else {
                number = number.toFixed(2);
            }
        }

        return number;
    };

    jQuery.fn.parseNumber.defaults = {
        locale: "us",
        decimalSeparatorAlwaysShown: false,
        isPercentage: false,            // treats the input as a percentage (i.e. input divided by 100)
        autoDetectPercentage: true,     // will search if the input string ends with '%', if it does then the above option is implicitly set
        isFullLocale: false,
        strict: false,                  // will abort the parse if it hits any unknown char
        overrideGroupSep: null,         // override for group separator
        overrideDecSep: null,           // override for decimal point separator
        overrideNegSign: null,          // override for negative sign
        allowPostfix: false             // will truncate the input string as soon as it hits an unknown char
    };

    jQuery.fn.formatNumber.defaults = {
        format: "#,###.00",
        locale: "us",
        decimalSeparatorAlwaysShown: false,
        nanForceZero: true,
        round: true,
        isFullLocale: false,
        overrideGroupSep: null,
        overrideDecSep: null,
        overrideNegSign: null,
        isPercentage: false,            // treats the input as a percentage (i.e. input multiplied by 100)
        autoDetectPercentage: true      // will search if the format string ends with '%', if it does then the above option is implicitly set
    };

    Number.prototype.toFixed = function (precision) {
        return jQuery._roundNumber(this, precision);
    };

    jQuery._roundNumber = function (number, decimalPlaces) {
        var power = Math.pow(10, decimalPlaces || 0);
        var value = String(Math.round(number * power) / power);

        // ensure the decimal places are there
        if (decimalPlaces > 0) {
            var dp = value.indexOf(".");
            if (dp == -1) {
                value += '.';
                dp = 0;
            } else {
                dp = value.length - (dp + 1);
            }

            while (dp < decimalPlaces) {
                value += '0';
                dp++;
            }
        }
        return value;
    };

})(jQuery);
!function(d){function g(){return new Date(Date.UTC.apply(Date,arguments))}function b(){var h=new Date();return g(h.getUTCFullYear(),h.getUTCMonth(),h.getUTCDate(),h.getUTCHours(),h.getUTCMinutes(),h.getUTCSeconds(),0)}var f=function(j,i){var l=this;this.element=d(j);this.container=i.container||"body";this.language=i.language||this.element.data("date-language")||"en";this.language=this.language in e?this.language:"en";this.isRTL=e[this.language].rtl||false;this.formatType=i.formatType||this.element.data("format-type")||"standard";this.format=c.parseFormat(i.format||this.element.data("date-format")||e[this.language].format||c.getDefaultFormat(this.formatType,"input"),this.formatType);this.isInline=false;this.isVisible=false;this.isInput=this.element.is("input");this.fontAwesome=i.fontAwesome||this.element.data("font-awesome")||false;this.bootcssVer=i.bootcssVer||(this.isInput?(this.element.is(".form-control")?3:2):(this.bootcssVer=this.element.is(".input-group")?3:2));this.component=this.element.is(".date")?(this.bootcssVer==3?this.element.find(".input-group-addon .glyphicon-th, .input-group-addon .glyphicon-time, .input-group-addon .glyphicon-calendar, .input-group-addon .glyphicon-calendar, .input-group-addon .fa-calendar, .input-group-addon .fa-clock-o").parent():this.element.find(".add-on .icon-th, .add-on .icon-time, .add-on .icon-calendar .fa-calendar .fa-clock-o").parent()):false;this.componentReset=this.element.is(".date")?(this.bootcssVer==3?this.element.find(".input-group-addon .glyphicon-remove .fa-times").parent():this.element.find(".add-on .icon-remove .fa-times").parent()):false;this.hasInput=this.component&&this.element.find("input").length;if(this.component&&this.component.length===0){this.component=false}this.linkField=i.linkField||this.element.data("link-field")||false;this.linkFormat=c.parseFormat(i.linkFormat||this.element.data("link-format")||c.getDefaultFormat(this.formatType,"link"),this.formatType);this.minuteStep=i.minuteStep||this.element.data("minute-step")||5;this.pickerPosition=i.pickerPosition||this.element.data("picker-position")||"bottom-right";this.showMeridian=i.showMeridian||this.element.data("show-meridian")||false;this.initialDate=i.initialDate||new Date();this.zIndex=i.zIndex||this.element.data("z-index")||undefined;this.icons={leftArrow:this.fontAwesome?"fa-arrow-left":(this.bootcssVer===3?"glyphicon-arrow-left":"icon-arrow-left"),rightArrow:this.fontAwesome?"fa-arrow-right":(this.bootcssVer===3?"glyphicon-arrow-right":"icon-arrow-right")};this.icontype=this.fontAwesome?"fa":"glyphicon";this._attachEvents();this.formatViewType="datetime";if("formatViewType" in i){this.formatViewType=i.formatViewType}else{if("formatViewType" in this.element.data()){this.formatViewType=this.element.data("formatViewType")}}this.minView=0;if("minView" in i){this.minView=i.minView}else{if("minView" in this.element.data()){this.minView=this.element.data("min-view")}}this.minView=c.convertViewMode(this.minView);this.maxView=c.modes.length-1;if("maxView" in i){this.maxView=i.maxView}else{if("maxView" in this.element.data()){this.maxView=this.element.data("max-view")}}this.maxView=c.convertViewMode(this.maxView);this.wheelViewModeNavigation=false;if("wheelViewModeNavigation" in i){this.wheelViewModeNavigation=i.wheelViewModeNavigation}else{if("wheelViewModeNavigation" in this.element.data()){this.wheelViewModeNavigation=this.element.data("view-mode-wheel-navigation")}}this.wheelViewModeNavigationInverseDirection=false;if("wheelViewModeNavigationInverseDirection" in i){this.wheelViewModeNavigationInverseDirection=i.wheelViewModeNavigationInverseDirection}else{if("wheelViewModeNavigationInverseDirection" in this.element.data()){this.wheelViewModeNavigationInverseDirection=this.element.data("view-mode-wheel-navigation-inverse-dir")}}this.wheelViewModeNavigationDelay=100;if("wheelViewModeNavigationDelay" in i){this.wheelViewModeNavigationDelay=i.wheelViewModeNavigationDelay}else{if("wheelViewModeNavigationDelay" in this.element.data()){this.wheelViewModeNavigationDelay=this.element.data("view-mode-wheel-navigation-delay")}}this.startViewMode=2;if("startView" in i){this.startViewMode=i.startView}else{if("startView" in this.element.data()){this.startViewMode=this.element.data("start-view")}}this.startViewMode=c.convertViewMode(this.startViewMode);this.viewMode=this.startViewMode;this.viewSelect=this.minView;if("viewSelect" in i){this.viewSelect=i.viewSelect}else{if("viewSelect" in this.element.data()){this.viewSelect=this.element.data("view-select")}}this.viewSelect=c.convertViewMode(this.viewSelect);this.forceParse=true;if("forceParse" in i){this.forceParse=i.forceParse}else{if("dateForceParse" in this.element.data()){this.forceParse=this.element.data("date-force-parse")}}var k=this.bootcssVer===3?c.templateV3:c.template;while(k.indexOf("{iconType}")!==-1){k=k.replace("{iconType}",this.icontype)}while(k.indexOf("{leftArrow}")!==-1){k=k.replace("{leftArrow}",this.icons.leftArrow)}while(k.indexOf("{rightArrow}")!==-1){k=k.replace("{rightArrow}",this.icons.rightArrow)}this.picker=d(k).appendTo(this.isInline?this.element:this.container).on({click:d.proxy(this.click,this),mousedown:d.proxy(this.mousedown,this)});if(this.wheelViewModeNavigation){if(d.fn.mousewheel){this.picker.on({mousewheel:d.proxy(this.mousewheel,this)})}else{console.log("Mouse Wheel event is not supported. Please include the jQuery Mouse Wheel plugin before enabling this option")}}if(this.isInline){this.picker.addClass("datetimepicker-inline")}else{this.picker.addClass("datetimepicker-dropdown-"+this.pickerPosition+" dropdown-menu")}if(this.isRTL){this.picker.addClass("datetimepicker-rtl");var h=this.bootcssVer===3?".prev span, .next span":".prev i, .next i";this.picker.find(h).toggleClass(this.icons.leftArrow+" "+this.icons.rightArrow)}d(document).on("mousedown",function(m){if(d(m.target).closest(".datetimepicker").length===0){l.hide()}});this.autoclose=false;if("autoclose" in i){this.autoclose=i.autoclose}else{if("dateAutoclose" in this.element.data()){this.autoclose=this.element.data("date-autoclose")}}this.keyboardNavigation=true;if("keyboardNavigation" in i){this.keyboardNavigation=i.keyboardNavigation}else{if("dateKeyboardNavigation" in this.element.data()){this.keyboardNavigation=this.element.data("date-keyboard-navigation")}}this.todayBtn=(i.todayBtn||this.element.data("date-today-btn")||false);this.todayHighlight=(i.todayHighlight||this.element.data("date-today-highlight")||false);this.weekStart=((i.weekStart||this.element.data("date-weekstart")||e[this.language].weekStart||0)%7);this.weekEnd=((this.weekStart+6)%7);this.startDate=-Infinity;this.endDate=Infinity;this.daysOfWeekDisabled=[];this.setStartDate(i.startDate||this.element.data("date-startdate"));this.setEndDate(i.endDate||this.element.data("date-enddate"));this.setDaysOfWeekDisabled(i.daysOfWeekDisabled||this.element.data("date-days-of-week-disabled"));this.setMinutesDisabled(i.minutesDisabled||this.element.data("date-minute-disabled"));this.setHoursDisabled(i.hoursDisabled||this.element.data("date-hour-disabled"));this.fillDow();this.fillMonths();this.update();this.showMode();if(this.isInline){this.show()}};f.prototype={constructor:f,_events:[],_attachEvents:function(){this._detachEvents();if(this.isInput){this._events=[[this.element,{focus:d.proxy(this.show,this),keyup:d.proxy(this.update,this),keydown:d.proxy(this.keydown,this)}]]}else{if(this.component&&this.hasInput){this._events=[[this.element.find("input"),{focus:d.proxy(this.show,this),keyup:d.proxy(this.update,this),keydown:d.proxy(this.keydown,this)}],[this.component,{click:d.proxy(this.show,this)}]];if(this.componentReset){this._events.push([this.componentReset,{click:d.proxy(this.reset,this)}])}}else{if(this.element.is("div")){this.isInline=true}else{this._events=[[this.element,{click:d.proxy(this.show,this)}]]}}}for(var h=0,j,k;h<this._events.length;h++){j=this._events[h][0];k=this._events[h][1];j.on(k)}},_detachEvents:function(){for(var h=0,j,k;h<this._events.length;h++){j=this._events[h][0];k=this._events[h][1];j.off(k)}this._events=[]},show:function(h){this.picker.show();this.height=this.component?this.component.outerHeight():this.element.outerHeight();if(this.forceParse){this.update()}this.place();d(window).on("resize",d.proxy(this.place,this));if(h){h.stopPropagation();h.preventDefault()}this.isVisible=true;this.element.trigger({type:"show",date:this.date})},hide:function(h){if(!this.isVisible){return}if(this.isInline){return}this.picker.hide();d(window).off("resize",this.place);this.viewMode=this.startViewMode;this.showMode();if(!this.isInput){d(document).off("mousedown",this.hide)}if(this.forceParse&&(this.isInput&&this.element.val()||this.hasInput&&this.element.find("input").val())){this.setValue()}this.isVisible=false;this.element.trigger({type:"hide",date:this.date})},remove:function(){this._detachEvents();this.picker.remove();delete this.picker;delete this.element.data().datetimepicker},getDate:function(){var h=this.getUTCDate();return new Date(h.getTime()+(h.getTimezoneOffset()*60000))},getUTCDate:function(){return this.date},setDate:function(h){this.setUTCDate(new Date(h.getTime()-(h.getTimezoneOffset()*60000)))},setUTCDate:function(h){if(h>=this.startDate&&h<=this.endDate){this.date=h;this.setValue();this.viewDate=this.date;this.fill()}else{this.element.trigger({type:"outOfRange",date:h,startDate:this.startDate,endDate:this.endDate})}},setFormat:function(i){this.format=c.parseFormat(i,this.formatType);var h;if(this.isInput){h=this.element}else{if(this.component){h=this.element.find("input")}}if(h&&h.val()){this.setValue()}},setValue:function(){var h=this.getFormattedDate();if(!this.isInput){if(this.component){this.element.find("input").val(h)}this.element.data("date",h)}else{this.element.val(h)}if(this.linkField){d("#"+this.linkField).val(this.getFormattedDate(this.linkFormat))}},getFormattedDate:function(h){if(h==undefined){h=this.format}return c.formatDate(this.date,h,this.language,this.formatType)},setStartDate:function(h){this.startDate=h||-Infinity;if(this.startDate!==-Infinity){this.startDate=c.parseDate(this.startDate,this.format,this.language,this.formatType)}this.update();this.updateNavArrows()},setEndDate:function(h){this.endDate=h||Infinity;if(this.endDate!==Infinity){this.endDate=c.parseDate(this.endDate,this.format,this.language,this.formatType)}this.update();this.updateNavArrows()},setDaysOfWeekDisabled:function(h){this.daysOfWeekDisabled=h||[];if(!d.isArray(this.daysOfWeekDisabled)){this.daysOfWeekDisabled=this.daysOfWeekDisabled.split(/,\s*/)}this.daysOfWeekDisabled=d.map(this.daysOfWeekDisabled,function(i){return parseInt(i,10)});this.update();this.updateNavArrows()},setMinutesDisabled:function(h){this.minutesDisabled=h||[];if(!d.isArray(this.minutesDisabled)){this.minutesDisabled=this.minutesDisabled.split(/,\s*/)}this.minutesDisabled=d.map(this.minutesDisabled,function(i){return parseInt(i,10)});this.update();this.updateNavArrows()},setHoursDisabled:function(h){this.hoursDisabled=h||[];if(!d.isArray(this.hoursDisabled)){this.hoursDisabled=this.hoursDisabled.split(/,\s*/)}this.hoursDisabled=d.map(this.hoursDisabled,function(i){return parseInt(i,10)});this.update();this.updateNavArrows()},place:function(){if(this.isInline){return}if(!this.zIndex){var h=0;d("div").each(function(){var m=parseInt(d(this).css("zIndex"),10);if(m>h){h=m}});this.zIndex=h+10}var l,k,j,i;if(this.container instanceof d){i=this.container.offset()}else{i=d(this.container).offset()}if(this.component){l=this.component.offset();j=l.left;if(this.pickerPosition=="bottom-left"||this.pickerPosition=="top-left"){j+=this.component.outerWidth()-this.picker.outerWidth()}}else{l=this.element.offset();j=l.left}if(j+220>document.body.clientWidth){j=document.body.clientWidth-220}if(this.pickerPosition=="top-left"||this.pickerPosition=="top-right"){k=l.top-this.picker.outerHeight()}else{k=l.top+this.height}k=k-i.top;j=j-i.left;k=k+document.body.scrollTop;this.picker.css({top:k,left:j,zIndex:this.zIndex})},update:function(){var h,i=false;if(arguments&&arguments.length&&(typeof arguments[0]==="string"||arguments[0] instanceof Date)){h=arguments[0];i=true}else{h=(this.isInput?this.element.val():this.element.find("input").val())||this.element.data("date")||this.initialDate;if(typeof h=="string"||h instanceof String){h=h.replace(/^\s+|\s+$/g,"")}}if(!h){h=new Date();i=false}this.date=c.parseDate(h,this.format,this.language,this.formatType);if(i){this.setValue()}if(this.date<this.startDate){this.viewDate=new Date(this.startDate)}else{if(this.date>this.endDate){this.viewDate=new Date(this.endDate)}else{this.viewDate=new Date(this.date)}}this.fill()},fillDow:function(){var h=this.weekStart,i="<tr>";while(h<this.weekStart+7){i+='<th class="dow">'+e[this.language].daysMin[(h++)%7]+"</th>"}i+="</tr>";this.picker.find(".datetimepicker-days thead").append(i)},fillMonths:function(){var j="",h=0;while(h<12){j+='<span class="month">'+e[this.language].monthsShort[h++]+"</span>"}this.picker.find(".datetimepicker-months td").html(j)},fill:function(){if(this.date==null||this.viewDate==null){return}var F=new Date(this.viewDate),s=F.getUTCFullYear(),G=F.getUTCMonth(),m=F.getUTCDate(),B=F.getUTCHours(),w=F.getUTCMinutes(),x=this.startDate!==-Infinity?this.startDate.getUTCFullYear():-Infinity,C=this.startDate!==-Infinity?this.startDate.getUTCMonth()+1:-Infinity,o=this.endDate!==Infinity?this.endDate.getUTCFullYear():Infinity,y=this.endDate!==Infinity?this.endDate.getUTCMonth()+1:Infinity,p=(new g(this.date.getUTCFullYear(),this.date.getUTCMonth(),this.date.getUTCDate())).valueOf(),E=new Date();this.picker.find(".datetimepicker-days thead th:eq(1)").text(e[this.language].months[G]+" "+s);if(this.formatViewType=="time"){var j=this.getFormattedDate();this.picker.find(".datetimepicker-hours thead th:eq(1)").text(j);this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(j)}else{this.picker.find(".datetimepicker-hours thead th:eq(1)").text(m+" "+e[this.language].months[G]+" "+s);this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(m+" "+e[this.language].months[G]+" "+s)}this.picker.find("tfoot th.today").text(e[this.language].today).toggle(this.todayBtn!==false);this.updateNavArrows();this.fillMonths();var I=g(s,G-1,28,0,0,0,0),A=c.getDaysInMonth(I.getUTCFullYear(),I.getUTCMonth());I.setUTCDate(A);I.setUTCDate(A-(I.getUTCDay()-this.weekStart+7)%7);var h=new Date(I);h.setUTCDate(h.getUTCDate()+42);h=h.valueOf();var q=[];var t;while(I.valueOf()<h){if(I.getUTCDay()==this.weekStart){q.push("<tr>")}t="";if(I.getUTCFullYear()<s||(I.getUTCFullYear()==s&&I.getUTCMonth()<G)){t+=" old"}else{if(I.getUTCFullYear()>s||(I.getUTCFullYear()==s&&I.getUTCMonth()>G)){t+=" new"}}if(this.todayHighlight&&I.getUTCFullYear()==E.getFullYear()&&I.getUTCMonth()==E.getMonth()&&I.getUTCDate()==E.getDate()){t+=" today"}if(I.valueOf()==p){t+=" active"}if((I.valueOf()+86400000)<=this.startDate||I.valueOf()>this.endDate||d.inArray(I.getUTCDay(),this.daysOfWeekDisabled)!==-1){t+=" disabled"}q.push('<td class="day'+t+'">'+I.getUTCDate()+"</td>");if(I.getUTCDay()==this.weekEnd){q.push("</tr>")}I.setUTCDate(I.getUTCDate()+1)}this.picker.find(".datetimepicker-days tbody").empty().append(q.join(""));q=[];var u="",D="",r="";var k=this.hoursDisabled||[];for(var z=0;z<24;z++){if(k.indexOf(z)!==-1){continue}var v=g(s,G,m,z);t="";if((v.valueOf()+3600000)<=this.startDate||v.valueOf()>this.endDate){t+=" disabled"}else{if(B==z){t+=" active"}}if(this.showMeridian&&e[this.language].meridiem.length==2){D=(z<12?e[this.language].meridiem[0]:e[this.language].meridiem[1]);if(D!=r){if(r!=""){q.push("</fieldset>")}q.push('<fieldset class="hour"><legend>'+D.toUpperCase()+"</legend>")}r=D;u=(z%12?z%12:12);q.push('<span class="hour'+t+" hour_"+(z<12?"am":"pm")+'">'+u+"</span>");if(z==23){q.push("</fieldset>")}}else{u=z+":00";q.push('<span class="hour'+t+'">'+u+"</span>")}}this.picker.find(".datetimepicker-hours td").html(q.join(""));q=[];u="",D="",r="";var l=this.minutesDisabled||[];for(var z=0;z<60;z+=this.minuteStep){if(l.indexOf(z)!==-1){continue}var v=g(s,G,m,B,z,0);t="";if(v.valueOf()<this.startDate||v.valueOf()>this.endDate){t+=" disabled"}else{if(Math.floor(w/this.minuteStep)==Math.floor(z/this.minuteStep)){t+=" active"}}if(this.showMeridian&&e[this.language].meridiem.length==2){D=(B<12?e[this.language].meridiem[0]:e[this.language].meridiem[1]);if(D!=r){if(r!=""){q.push("</fieldset>")}q.push('<fieldset class="minute"><legend>'+D.toUpperCase()+"</legend>")}r=D;u=(B%12?B%12:12);q.push('<span class="minute'+t+'">'+u+":"+(z<10?"0"+z:z)+"</span>");if(z==59){q.push("</fieldset>")}}else{u=z+":00";q.push('<span class="minute'+t+'">'+B+":"+(z<10?"0"+z:z)+"</span>")}}this.picker.find(".datetimepicker-minutes td").html(q.join(""));var J=this.date.getUTCFullYear();var n=this.picker.find(".datetimepicker-months").find("th:eq(1)").text(s).end().find("span").removeClass("active");if(J==s){n.eq(this.date.getUTCMonth()+2).addClass("active")}if(s<x||s>o){n.addClass("disabled")}if(s==x){n.slice(0,C+1).addClass("disabled")}if(s==o){n.slice(y).addClass("disabled")}q="";s=parseInt(s/10,10)*10;var H=this.picker.find(".datetimepicker-years").find("th:eq(1)").text(s+"-"+(s+9)).end().find("td");s-=1;for(var z=-1;z<11;z++){q+='<span class="year'+(z==-1||z==10?" old":"")+(J==s?" active":"")+(s<x||s>o?" disabled":"")+'">'+s+"</span>";s+=1}H.html(q);this.place()},updateNavArrows:function(){var l=new Date(this.viewDate),j=l.getUTCFullYear(),k=l.getUTCMonth(),i=l.getUTCDate(),h=l.getUTCHours();switch(this.viewMode){case 0:if(this.startDate!==-Infinity&&j<=this.startDate.getUTCFullYear()&&k<=this.startDate.getUTCMonth()&&i<=this.startDate.getUTCDate()&&h<=this.startDate.getUTCHours()){this.picker.find(".prev").css({visibility:"hidden"})}else{this.picker.find(".prev").css({visibility:"visible"})}if(this.endDate!==Infinity&&j>=this.endDate.getUTCFullYear()&&k>=this.endDate.getUTCMonth()&&i>=this.endDate.getUTCDate()&&h>=this.endDate.getUTCHours()){this.picker.find(".next").css({visibility:"hidden"})}else{this.picker.find(".next").css({visibility:"visible"})}break;case 1:if(this.startDate!==-Infinity&&j<=this.startDate.getUTCFullYear()&&k<=this.startDate.getUTCMonth()&&i<=this.startDate.getUTCDate()){this.picker.find(".prev").css({visibility:"hidden"})}else{this.picker.find(".prev").css({visibility:"visible"})}if(this.endDate!==Infinity&&j>=this.endDate.getUTCFullYear()&&k>=this.endDate.getUTCMonth()&&i>=this.endDate.getUTCDate()){this.picker.find(".next").css({visibility:"hidden"})}else{this.picker.find(".next").css({visibility:"visible"})}break;case 2:if(this.startDate!==-Infinity&&j<=this.startDate.getUTCFullYear()&&k<=this.startDate.getUTCMonth()){this.picker.find(".prev").css({visibility:"hidden"})}else{this.picker.find(".prev").css({visibility:"visible"})}if(this.endDate!==Infinity&&j>=this.endDate.getUTCFullYear()&&k>=this.endDate.getUTCMonth()){this.picker.find(".next").css({visibility:"hidden"})}else{this.picker.find(".next").css({visibility:"visible"})}break;case 3:case 4:if(this.startDate!==-Infinity&&j<=this.startDate.getUTCFullYear()){this.picker.find(".prev").css({visibility:"hidden"})}else{this.picker.find(".prev").css({visibility:"visible"})}if(this.endDate!==Infinity&&j>=this.endDate.getUTCFullYear()){this.picker.find(".next").css({visibility:"hidden"})}else{this.picker.find(".next").css({visibility:"visible"})}break}},mousewheel:function(i){i.preventDefault();i.stopPropagation();if(this.wheelPause){return}this.wheelPause=true;var h=i.originalEvent;var k=h.wheelDelta;var j=k>0?1:(k===0)?0:-1;if(this.wheelViewModeNavigationInverseDirection){j=-j}this.showMode(j);setTimeout(d.proxy(function(){this.wheelPause=false},this),this.wheelViewModeNavigationDelay)},click:function(l){l.stopPropagation();l.preventDefault();var m=d(l.target).closest("span, td, th, legend");if(m.is("."+this.icontype)){m=d(m).parent().closest("span, td, th, legend")}if(m.length==1){if(m.is(".disabled")){this.element.trigger({type:"outOfRange",date:this.viewDate,startDate:this.startDate,endDate:this.endDate});return}switch(m[0].nodeName.toLowerCase()){case"th":switch(m[0].className){case"switch":this.showMode(1);break;case"prev":case"next":var h=c.modes[this.viewMode].navStep*(m[0].className=="prev"?-1:1);switch(this.viewMode){case 0:this.viewDate=this.moveHour(this.viewDate,h);break;case 1:this.viewDate=this.moveDate(this.viewDate,h);break;case 2:this.viewDate=this.moveMonth(this.viewDate,h);break;case 3:case 4:this.viewDate=this.moveYear(this.viewDate,h);break}this.fill();this.element.trigger({type:m[0].className+":"+this.convertViewModeText(this.viewMode),date:this.viewDate,startDate:this.startDate,endDate:this.endDate});break;case"today":var i=new Date();i=g(i.getFullYear(),i.getMonth(),i.getDate(),i.getHours(),i.getMinutes(),i.getSeconds(),0);if(i<this.startDate){i=this.startDate}else{if(i>this.endDate){i=this.endDate}}this.viewMode=this.startViewMode;this.showMode(0);this._setDate(i);this.fill();if(this.autoclose){this.hide()}break}break;case"span":if(!m.is(".disabled")){var o=this.viewDate.getUTCFullYear(),n=this.viewDate.getUTCMonth(),p=this.viewDate.getUTCDate(),q=this.viewDate.getUTCHours(),j=this.viewDate.getUTCMinutes(),r=this.viewDate.getUTCSeconds();if(m.is(".month")){this.viewDate.setUTCDate(1);n=m.parent().find("span").index(m);p=this.viewDate.getUTCDate();this.viewDate.setUTCMonth(n);this.element.trigger({type:"changeMonth",date:this.viewDate});if(this.viewSelect>=3){this._setDate(g(o,n,p,q,j,r,0))}}else{if(m.is(".year")){this.viewDate.setUTCDate(1);o=parseInt(m.text(),10)||0;this.viewDate.setUTCFullYear(o);this.element.trigger({type:"changeYear",date:this.viewDate});if(this.viewSelect>=4){this._setDate(g(o,n,p,q,j,r,0))}}else{if(m.is(".hour")){q=parseInt(m.text(),10)||0;if(m.hasClass("hour_am")||m.hasClass("hour_pm")){if(q==12&&m.hasClass("hour_am")){q=0}else{if(q!=12&&m.hasClass("hour_pm")){q+=12}}}this.viewDate.setUTCHours(q);this.element.trigger({type:"changeHour",date:this.viewDate});if(this.viewSelect>=1){this._setDate(g(o,n,p,q,j,r,0))}}else{if(m.is(".minute")){j=parseInt(m.text().substr(m.text().indexOf(":")+1),10)||0;this.viewDate.setUTCMinutes(j);this.element.trigger({type:"changeMinute",date:this.viewDate});if(this.viewSelect>=0){this._setDate(g(o,n,p,q,j,r,0))}}}}}if(this.viewMode!=0){var k=this.viewMode;this.showMode(-1);this.fill();if(k==this.viewMode&&this.autoclose){this.hide()}}else{this.fill();if(this.autoclose){this.hide()}}}break;case"td":if(m.is(".day")&&!m.is(".disabled")){var p=parseInt(m.text(),10)||1;var o=this.viewDate.getUTCFullYear(),n=this.viewDate.getUTCMonth(),q=this.viewDate.getUTCHours(),j=this.viewDate.getUTCMinutes(),r=this.viewDate.getUTCSeconds();if(m.is(".old")){if(n===0){n=11;o-=1}else{n-=1}}else{if(m.is(".new")){if(n==11){n=0;o+=1}else{n+=1}}}this.viewDate.setUTCFullYear(o);this.viewDate.setUTCMonth(n,p);this.element.trigger({type:"changeDay",date:this.viewDate});if(this.viewSelect>=2){this._setDate(g(o,n,p,q,j,r,0))}}var k=this.viewMode;this.showMode(-1);this.fill();if(k==this.viewMode&&this.autoclose){this.hide()}break}}},_setDate:function(h,j){if(!j||j=="date"){this.date=h}if(!j||j=="view"){this.viewDate=h}this.fill();this.setValue();var i;if(this.isInput){i=this.element}else{if(this.component){i=this.element.find("input")}}if(i){i.change();if(this.autoclose&&(!j||j=="date")){}}this.element.trigger({type:"changeDate",date:this.date})},moveMinute:function(i,h){if(!h){return i}var j=new Date(i.valueOf());j.setUTCMinutes(j.getUTCMinutes()+(h*this.minuteStep));return j},moveHour:function(i,h){if(!h){return i}var j=new Date(i.valueOf());j.setUTCHours(j.getUTCHours()+h);return j},moveDate:function(i,h){if(!h){return i}var j=new Date(i.valueOf());j.setUTCDate(j.getUTCDate()+h);return j},moveMonth:function(h,j){if(!j){return h}var m=new Date(h.valueOf()),q=m.getUTCDate(),n=m.getUTCMonth(),l=Math.abs(j),p,o;j=j>0?1:-1;if(l==1){o=j==-1?function(){return m.getUTCMonth()==n}:function(){return m.getUTCMonth()!=p};p=n+j;m.setUTCMonth(p);if(p<0||p>11){p=(p+12)%12}}else{for(var k=0;k<l;k++){m=this.moveMonth(m,j)}p=m.getUTCMonth();m.setUTCDate(q);o=function(){return p!=m.getUTCMonth()}}while(o()){m.setUTCDate(--q);m.setUTCMonth(p)}return m},moveYear:function(i,h){return this.moveMonth(i,h*12)},dateWithinRange:function(h){return h>=this.startDate&&h<=this.endDate},keydown:function(l){if(this.picker.is(":not(:visible)")){if(l.keyCode==27){this.show()}return}var n=false,i,o,m,p,h;switch(l.keyCode){case 27:this.hide();l.preventDefault();break;case 37:case 39:if(!this.keyboardNavigation){break}i=l.keyCode==37?-1:1;viewMode=this.viewMode;if(l.ctrlKey){viewMode+=2}else{if(l.shiftKey){viewMode+=1}}if(viewMode==4){p=this.moveYear(this.date,i);h=this.moveYear(this.viewDate,i)}else{if(viewMode==3){p=this.moveMonth(this.date,i);h=this.moveMonth(this.viewDate,i)}else{if(viewMode==2){p=this.moveDate(this.date,i);h=this.moveDate(this.viewDate,i)}else{if(viewMode==1){p=this.moveHour(this.date,i);h=this.moveHour(this.viewDate,i)}else{if(viewMode==0){p=this.moveMinute(this.date,i);h=this.moveMinute(this.viewDate,i)}}}}}if(this.dateWithinRange(p)){this.date=p;this.viewDate=h;this.setValue();this.update();l.preventDefault();n=true}break;case 38:case 40:if(!this.keyboardNavigation){break}i=l.keyCode==38?-1:1;viewMode=this.viewMode;if(l.ctrlKey){viewMode+=2}else{if(l.shiftKey){viewMode+=1}}if(viewMode==4){p=this.moveYear(this.date,i);h=this.moveYear(this.viewDate,i)}else{if(viewMode==3){p=this.moveMonth(this.date,i);h=this.moveMonth(this.viewDate,i)}else{if(viewMode==2){p=this.moveDate(this.date,i*7);h=this.moveDate(this.viewDate,i*7)}else{if(viewMode==1){if(this.showMeridian){p=this.moveHour(this.date,i*6);h=this.moveHour(this.viewDate,i*6)}else{p=this.moveHour(this.date,i*4);h=this.moveHour(this.viewDate,i*4)}}else{if(viewMode==0){p=this.moveMinute(this.date,i*4);h=this.moveMinute(this.viewDate,i*4)}}}}}if(this.dateWithinRange(p)){this.date=p;this.viewDate=h;this.setValue();this.update();l.preventDefault();n=true}break;case 13:if(this.viewMode!=0){var k=this.viewMode;this.showMode(-1);this.fill();if(k==this.viewMode&&this.autoclose){this.hide()}}else{this.fill();if(this.autoclose){this.hide()}}l.preventDefault();break;case 9:this.hide();break}if(n){var j;if(this.isInput){j=this.element}else{if(this.component){j=this.element.find("input")}}if(j){j.change()}this.element.trigger({type:"changeDate",date:this.date})}},showMode:function(h){if(h){var i=Math.max(0,Math.min(c.modes.length-1,this.viewMode+h));if(i>=this.minView&&i<=this.maxView){this.element.trigger({type:"changeMode",date:this.viewDate,oldViewMode:this.viewMode,newViewMode:i});this.viewMode=i}}this.picker.find(">div").hide().filter(".datetimepicker-"+c.modes[this.viewMode].clsName).css("display","block");this.updateNavArrows()},reset:function(h){this._setDate(null,"date")},convertViewModeText:function(h){switch(h){case 4:return"decade";case 3:return"year";case 2:return"month";case 1:return"day";case 0:return"hour"}}};var a=d.fn.datetimepicker;d.fn.datetimepicker=function(j){var h=Array.apply(null,arguments);h.shift();var i;this.each(function(){var m=d(this),l=m.data("datetimepicker"),k=typeof j=="object"&&j;if(!l){m.data("datetimepicker",(l=new f(this,d.extend({},d.fn.datetimepicker.defaults,k))))}if(typeof j=="string"&&typeof l[j]=="function"){i=l[j].apply(l,h);if(i!==undefined){return false}}});if(i!==undefined){return i}else{return this}};d.fn.datetimepicker.defaults={};d.fn.datetimepicker.Constructor=f;var e=d.fn.datetimepicker.dates={en:{days:["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"],daysShort:["Sun","Mon","Tue","Wed","Thu","Fri","Sat","Sun"],daysMin:["Su","Mo","Tu","We","Th","Fr","Sa","Su"],months:["January","February","March","April","May","June","July","August","September","October","November","December"],monthsShort:["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"],meridiem:["am","pm"],suffix:["st","nd","rd","th"],today:"Today"}};var c={modes:[{clsName:"minutes",navFnc:"Hours",navStep:1},{clsName:"hours",navFnc:"Date",navStep:1},{clsName:"days",navFnc:"Month",navStep:1},{clsName:"months",navFnc:"FullYear",navStep:1},{clsName:"years",navFnc:"FullYear",navStep:10}],isLeapYear:function(h){return(((h%4===0)&&(h%100!==0))||(h%400===0))},getDaysInMonth:function(h,i){return[31,(c.isLeapYear(h)?29:28),31,30,31,30,31,31,30,31,30,31][i]},getDefaultFormat:function(h,i){if(h=="standard"){if(i=="input"){return"yyyy-mm-dd hh:ii"}else{return"yyyy-mm-dd hh:ii:ss"}}else{if(h=="php"){if(i=="input"){return"Y-m-d H:i"}else{return"Y-m-d H:i:s"}}else{throw new Error("Invalid format type.")}}},validParts:function(h){if(h=="standard"){return/hh?|HH?|p|P|ii?|ss?|dd?|DD?|mm?|MM?|yy(?:yy)?/g}else{if(h=="php"){return/[dDjlNwzFmMnStyYaABgGhHis]/g}else{throw new Error("Invalid format type.")}}},nonpunctuation:/[^ -\/:-@\[-`{-~\t\n\rTZ]+/g,parseFormat:function(k,i){var h=k.replace(this.validParts(i),"\0").split("\0"),j=k.match(this.validParts(i));if(!h||!h.length||!j||j.length==0){throw new Error("Invalid date format.")}return{separators:h,parts:j}},parseDate:function(m,v,p,t){if(m instanceof Date){var x=new Date(m.valueOf()-m.getTimezoneOffset()*60000);x.setMilliseconds(0);return x}if(/^\d{4}\-\d{1,2}\-\d{1,2}$/.test(m)){v=this.parseFormat("yyyy-mm-dd",t)}if(/^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}$/.test(m)){v=this.parseFormat("yyyy-mm-dd hh:ii",t)}if(/^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}\:\d{1,2}[Z]{0,1}$/.test(m)){v=this.parseFormat("yyyy-mm-dd hh:ii:ss",t)}if(/^[-+]\d+[dmwy]([\s,]+[-+]\d+[dmwy])*$/.test(m)){var y=/([-+]\d+)([dmwy])/,n=m.match(/([-+]\d+)([dmwy])/g),h,l;m=new Date();for(var o=0;o<n.length;o++){h=y.exec(n[o]);l=parseInt(h[1]);switch(h[2]){case"d":m.setUTCDate(m.getUTCDate()+l);break;case"m":m=f.prototype.moveMonth.call(f.prototype,m,l);break;case"w":m.setUTCDate(m.getUTCDate()+l*7);break;case"y":m=f.prototype.moveYear.call(f.prototype,m,l);break}}return g(m.getUTCFullYear(),m.getUTCMonth(),m.getUTCDate(),m.getUTCHours(),m.getUTCMinutes(),m.getUTCSeconds(),0)}var n=m&&m.toString().match(this.nonpunctuation)||[],m=new Date(0,0,0,0,0,0,0),r={},u=["hh","h","ii","i","ss","s","yyyy","yy","M","MM","m","mm","D","DD","d","dd","H","HH","p","P"],w={hh:function(s,i){return s.setUTCHours(i)},h:function(s,i){return s.setUTCHours(i)},HH:function(s,i){return s.setUTCHours(i==12?0:i)},H:function(s,i){return s.setUTCHours(i==12?0:i)},ii:function(s,i){return s.setUTCMinutes(i)},i:function(s,i){return s.setUTCMinutes(i)},ss:function(s,i){return s.setUTCSeconds(i)},s:function(s,i){return s.setUTCSeconds(i)},yyyy:function(s,i){return s.setUTCFullYear(i)},yy:function(s,i){return s.setUTCFullYear(2000+i)},m:function(s,i){i-=1;while(i<0){i+=12}i%=12;s.setUTCMonth(i);while(s.getUTCMonth()!=i){if(isNaN(s.getUTCMonth())){return s}else{s.setUTCDate(s.getUTCDate()-1)}}return s},d:function(s,i){return s.setUTCDate(i)},p:function(s,i){return s.setUTCHours(i==1?s.getUTCHours()+12:s.getUTCHours())}},k,q,h;w.M=w.MM=w.mm=w.m;w.dd=w.d;w.P=w.p;m=g(m.getFullYear(),m.getMonth(),m.getDate(),m.getHours(),m.getMinutes(),m.getSeconds());if(n.length==v.parts.length){for(var o=0,j=v.parts.length;o<j;o++){k=parseInt(n[o],10);h=v.parts[o];if(isNaN(k)){switch(h){case"MM":q=d(e[p].months).filter(function(){var i=this.slice(0,n[o].length),s=n[o].slice(0,i.length);return i==s});k=d.inArray(q[0],e[p].months)+1;break;case"M":q=d(e[p].monthsShort).filter(function(){var i=this.slice(0,n[o].length),s=n[o].slice(0,i.length);return i.toLowerCase()==s.toLowerCase()});k=d.inArray(q[0],e[p].monthsShort)+1;break;case"p":case"P":k=d.inArray(n[o].toLowerCase(),e[p].meridiem);break}}r[h]=k}for(var o=0,z;o<u.length;o++){z=u[o];if(z in r&&!isNaN(r[z])){w[z](m,r[z])}}}return m},formatDate:function(h,n,p,l){if(h==null){return""}var o;if(l=="standard"){o={yy:h.getUTCFullYear().toString().substring(2),yyyy:h.getUTCFullYear(),m:h.getUTCMonth()+1,M:e[p].monthsShort[h.getUTCMonth()],MM:e[p].months[h.getUTCMonth()],d:h.getUTCDate(),D:e[p].daysShort[h.getUTCDay()],DD:e[p].days[h.getUTCDay()],p:(e[p].meridiem.length==2?e[p].meridiem[h.getUTCHours()<12?0:1]:""),h:h.getUTCHours(),i:h.getUTCMinutes(),s:h.getUTCSeconds()};if(e[p].meridiem.length==2){o.H=(o.h%12==0?12:o.h%12)}else{o.H=o.h}o.HH=(o.H<10?"0":"")+o.H;o.P=o.p.toUpperCase();o.hh=(o.h<10?"0":"")+o.h;o.ii=(o.i<10?"0":"")+o.i;o.ss=(o.s<10?"0":"")+o.s;o.dd=(o.d<10?"0":"")+o.d;o.mm=(o.m<10?"0":"")+o.m}else{if(l=="php"){o={y:h.getUTCFullYear().toString().substring(2),Y:h.getUTCFullYear(),F:e[p].months[h.getUTCMonth()],M:e[p].monthsShort[h.getUTCMonth()],n:h.getUTCMonth()+1,t:c.getDaysInMonth(h.getUTCFullYear(),h.getUTCMonth()),j:h.getUTCDate(),l:e[p].days[h.getUTCDay()],D:e[p].daysShort[h.getUTCDay()],w:h.getUTCDay(),N:(h.getUTCDay()==0?7:h.getUTCDay()),S:(h.getUTCDate()%10<=e[p].suffix.length?e[p].suffix[h.getUTCDate()%10-1]:""),a:(e[p].meridiem.length==2?e[p].meridiem[h.getUTCHours()<12?0:1]:""),g:(h.getUTCHours()%12==0?12:h.getUTCHours()%12),G:h.getUTCHours(),i:h.getUTCMinutes(),s:h.getUTCSeconds()};o.m=(o.n<10?"0":"")+o.n;o.d=(o.j<10?"0":"")+o.j;o.A=o.a.toString().toUpperCase();o.h=(o.g<10?"0":"")+o.g;o.H=(o.G<10?"0":"")+o.G;o.i=(o.i<10?"0":"")+o.i;o.s=(o.s<10?"0":"")+o.s}else{throw new Error("Invalid format type.")}}var h=[],m=d.extend([],n.separators);for(var k=0,j=n.parts.length;k<j;k++){if(m.length){h.push(m.shift())}h.push(o[n.parts[k]])}if(m.length){h.push(m.shift())}return h.join("")},convertViewMode:function(h){switch(h){case 4:case"decade":h=4;break;case 3:case"year":h=3;break;case 2:case"month":h=2;break;case 1:case"day":h=1;break;case 0:case"hour":h=0;break}return h},headTemplate:'<thead><tr><th class="prev"><i class="{leftArrow}"/></th><th colspan="5" class="switch"></th><th class="next"><i class="{rightArrow}"/></th></tr></thead>',headTemplateV3:'<thead><tr><th class="prev"><span class="{iconType} {leftArrow}"></span> </th><th colspan="5" class="switch"></th><th class="next"><span class="{iconType} {rightArrow}"></span> </th></tr></thead>',contTemplate:'<tbody><tr><td colspan="7"></td></tr></tbody>',footTemplate:'<tfoot><tr><th colspan="7" class="today"></th></tr></tfoot>'};c.template='<div class="datetimepicker"><div class="datetimepicker-minutes"><table class=" table-condensed">'+c.headTemplate+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-hours"><table class=" table-condensed">'+c.headTemplate+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-days"><table class=" table-condensed">'+c.headTemplate+"<tbody></tbody>"+c.footTemplate+'</table></div><div class="datetimepicker-months"><table class="table-condensed">'+c.headTemplate+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-years"><table class="table-condensed">'+c.headTemplate+c.contTemplate+c.footTemplate+"</table></div></div>";c.templateV3='<div class="datetimepicker"><div class="datetimepicker-minutes"><table class=" table-condensed">'+c.headTemplateV3+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-hours"><table class=" table-condensed">'+c.headTemplateV3+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-days"><table class=" table-condensed">'+c.headTemplateV3+"<tbody></tbody>"+c.footTemplate+'</table></div><div class="datetimepicker-months"><table class="table-condensed">'+c.headTemplateV3+c.contTemplate+c.footTemplate+'</table></div><div class="datetimepicker-years"><table class="table-condensed">'+c.headTemplateV3+c.contTemplate+c.footTemplate+"</table></div></div>";d.fn.datetimepicker.DPGlobal=c;d.fn.datetimepicker.noConflict=function(){d.fn.datetimepicker=a;return this};d(document).on("focus.datetimepicker.data-api click.datetimepicker.data-api",'[data-provide="datetimepicker"]',function(i){var h=d(this);if(h.data("datetimepicker")){return}i.preventDefault();h.datetimepicker("show")});d(function(){d('[data-provide="datetimepicker-inline"]').datetimepicker()})}(window.jQuery);
/**
 * Simplified Chinese translation for bootstrap-datetimepicker
 * Yuan Cheung <advanimal@gmail.com>
 */
;(function($){
	$.fn.datetimepicker.dates['zh-CN'] = {
			days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
			daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
			daysMin:  ["日", "一", "二", "三", "四", "五", "六", "日"],
			months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
			monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
			today: "今天",
			suffix: [],
			meridiem: ["上午", "下午"]
	};
}(jQuery));

!function(a){"use strict";function b(a,b){for(var c=0;c<a.length;++c)b(a[c],c)}function c(b,c){this.$select=a(b),this.$select.attr("data-placeholder")&&(c.nonSelectedText=this.$select.data("placeholder")),this.options=this.mergeOptions(a.extend({},c,this.$select.data())),this.originalOptions=this.$select.clone()[0].options,this.query="",this.searchTimeout=null,this.lastToggledInput=null,this.options.multiple="multiple"===this.$select.attr("multiple"),this.options.onChange=a.proxy(this.options.onChange,this),this.options.onDropdownShow=a.proxy(this.options.onDropdownShow,this),this.options.onDropdownHide=a.proxy(this.options.onDropdownHide,this),this.options.onDropdownShown=a.proxy(this.options.onDropdownShown,this),this.options.onDropdownHidden=a.proxy(this.options.onDropdownHidden,this),this.buildContainer(),this.buildButton(),this.buildDropdown(),this.buildSelectAll(),this.buildDropdownOptions(),this.buildFilter(),this.updateButtonText(),this.updateSelectAll(),this.options.disableIfEmpty&&a("option",this.$select).length<=0&&this.disable(),this.$select.hide().after(this.$container)}"undefined"!=typeof ko&&ko.bindingHandlers&&!ko.bindingHandlers.multiselect&&(ko.bindingHandlers.multiselect={after:["options","value","selectedOptions"],init:function(b,c,d,e,f){var g=a(b),h=ko.toJS(c());if(g.multiselect(h),d.has("options")){var i=d.get("options");ko.isObservable(i)&&ko.computed({read:function(){i(),setTimeout(function(){var a=g.data("multiselect");a&&a.updateOriginalOptions(),g.multiselect("rebuild")},1)},disposeWhenNodeIsRemoved:b})}if(d.has("value")){var j=d.get("value");ko.isObservable(j)&&ko.computed({read:function(){j(),setTimeout(function(){g.multiselect("refresh")},1)},disposeWhenNodeIsRemoved:b}).extend({rateLimit:100,notifyWhenChangesStop:!0})}if(d.has("selectedOptions")){var k=d.get("selectedOptions");ko.isObservable(k)&&ko.computed({read:function(){k(),setTimeout(function(){g.multiselect("refresh")},1)},disposeWhenNodeIsRemoved:b}).extend({rateLimit:100,notifyWhenChangesStop:!0})}ko.utils.domNodeDisposal.addDisposeCallback(b,function(){g.multiselect("destroy")})},update:function(b,c,d,e,f){var g=a(b),h=ko.toJS(c());g.multiselect("setOptions",h),g.multiselect("rebuild")}}),c.prototype={defaults:{buttonText:function(b,c){if(0===b.length)return this.nonSelectedText;if(this.allSelectedText&&b.length===a("option",a(c)).length&&1!==a("option",a(c)).length&&this.multiple)return this.selectAllNumber?this.allSelectedText+" ("+b.length+")":this.allSelectedText;if(b.length>this.numberDisplayed)return b.length+" "+this.nSelectedText;var d="",e=this.delimiterText;return b.each(function(){var b=void 0!==a(this).attr("label")?a(this).attr("label"):a(this).text();d+=b+e}),d.substr(0,d.length-2)},buttonTitle:function(b,c){if(0===b.length)return this.nonSelectedText;var d="",e=this.delimiterText;return b.each(function(){var b=void 0!==a(this).attr("label")?a(this).attr("label"):a(this).text();d+=b+e}),d.substr(0,d.length-2)},optionLabel:function(b){return a(b).attr("label")||a(b).text()},onChange:function(a,b){},onDropdownShow:function(a){},onDropdownHide:function(a){},onDropdownShown:function(a){},onDropdownHidden:function(a){},onSelectAll:function(){},enableHTML:!1,buttonClass:"btn btn-default",inheritClass:!1,buttonWidth:"auto",buttonContainer:'<div class="btn-group" />',dropRight:!1,selectedClass:"active",maxHeight:!1,checkboxName:!1,includeSelectAllOption:!1,includeSelectAllIfMoreThan:0,selectAllText:" Select all",selectAllValue:"multiselect-all",selectAllName:!1,selectAllNumber:!0,enableFiltering:!1,enableCaseInsensitiveFiltering:!1,enableClickableOptGroups:!1,filterPlaceholder:"Search",filterBehavior:"text",includeFilterClearBtn:!0,preventInputChangeEvent:!1,nonSelectedText:"None selected",nSelectedText:"selected",allSelectedText:"All selected",numberDisplayed:3,disableIfEmpty:!1,delimiterText:", ",templates:{button:'<button type="button" class="multiselect dropdown-toggle" data-toggle="dropdown"><span class="multiselect-selected-text"></span> <b class="caret"></b></button>',ul:'<ul class="multiselect-container dropdown-menu"></ul>',filter:'<li class="multiselect-item filter"><div class="input-group"><span class="input-group-addon"><i class="glyphicon glyphicon-search"></i></span><input class="form-control multiselect-search" type="text"></div></li>',filterClearBtn:'<span class="input-group-btn"><button class="btn btn-default multiselect-clear-filter" type="button"><i class="glyphicon glyphicon-remove-circle"></i></button></span>',li:'<li><a tabindex="0"><label></label></a></li>',divider:'<li class="multiselect-item divider"></li>',liGroup:'<li class="multiselect-item multiselect-group"><label></label></li>'}},constructor:c,buildContainer:function(){this.$container=a(this.options.buttonContainer),this.$container.on("show.bs.dropdown",this.options.onDropdownShow),this.$container.on("hide.bs.dropdown",this.options.onDropdownHide),this.$container.on("shown.bs.dropdown",this.options.onDropdownShown),this.$container.on("hidden.bs.dropdown",this.options.onDropdownHidden)},buildButton:function(){this.$button=a(this.options.templates.button).addClass(this.options.buttonClass),this.$select.attr("class")&&this.options.inheritClass&&this.$button.addClass(this.$select.attr("class")),this.$select.prop("disabled")?this.disable():this.enable(),this.options.buttonWidth&&"auto"!==this.options.buttonWidth&&(this.$button.css({width:this.options.buttonWidth,overflow:"hidden","text-overflow":"ellipsis"}),this.$container.css({width:this.options.buttonWidth}));var b=this.$select.attr("tabindex");b&&this.$button.attr("tabindex",b),this.$container.prepend(this.$button)},buildDropdown:function(){this.$ul=a(this.options.templates.ul),this.options.dropRight&&this.$ul.addClass("pull-right"),this.options.maxHeight&&this.$ul.css({"max-height":this.options.maxHeight+"px","overflow-y":"auto","overflow-x":"hidden"}),this.$container.append(this.$ul)},buildDropdownOptions:function(){this.$select.children().each(a.proxy(function(b,c){var d=a(c),e=d.prop("tagName").toLowerCase();d.prop("value")!==this.options.selectAllValue&&("optgroup"===e?this.createOptgroup(c):"option"===e&&("divider"===d.data("role")?this.createDivider():this.createOptionValue(c)))},this)),a("li input",this.$ul).on("change",a.proxy(function(b){var c=a(b.target),d=c.prop("checked")||!1,e=c.val()===this.options.selectAllValue;this.options.selectedClass&&(d?c.closest("li").addClass(this.options.selectedClass):c.closest("li").removeClass(this.options.selectedClass));var f=c.val(),g=this.getOptionByValue(f),h=a("option",this.$select).not(g),i=a("input",this.$container).not(c);return e&&(d?this.selectAll():this.deselectAll()),e||(d?(g.prop("selected",!0),this.options.multiple?g.prop("selected",!0):(this.options.selectedClass&&a(i).closest("li").removeClass(this.options.selectedClass),a(i).prop("checked",!1),h.prop("selected",!1),this.$button.click()),"active"===this.options.selectedClass&&h.closest("a").css("outline","")):g.prop("selected",!1)),this.$select.change(),this.updateButtonText(),this.updateSelectAll(),this.options.onChange(g,d),this.options.preventInputChangeEvent?!1:void 0},this)),a("li a",this.$ul).on("mousedown",function(a){return a.shiftKey?!1:void 0}),a("li a",this.$ul).on("touchstart click",a.proxy(function(b){b.stopPropagation();var c=a(b.target);if(b.shiftKey&&this.options.multiple){c.is("label")&&(b.preventDefault(),c=c.find("input"),c.prop("checked",!c.prop("checked")));var d=c.prop("checked")||!1;if(null!==this.lastToggledInput&&this.lastToggledInput!==c){var e=c.closest("li").index(),f=this.lastToggledInput.closest("li").index();if(e>f){var g=f;f=e,e=g}++f;var h=this.$ul.find("li").slice(e,f).find("input");h.prop("checked",d),this.options.selectedClass&&h.closest("li").toggleClass(this.options.selectedClass,d);for(var i=0,j=h.length;j>i;i++){var k=a(h[i]),l=this.getOptionByValue(k.val());l.prop("selected",d)}}c.trigger("change")}c.is("input")&&!c.closest("li").is(".multiselect-item")&&(this.lastToggledInput=c),c.blur()},this)),this.$container.off("keydown.multiselect").on("keydown.multiselect",a.proxy(function(b){if(!a('input[type="text"]',this.$container).is(":focus"))if(9===b.keyCode&&this.$container.hasClass("open"))this.$button.click();else{var c=a(this.$container).find("li:not(.divider):not(.disabled) a").filter(":visible");if(!c.length)return;var d=c.index(c.filter(":focus"));38===b.keyCode&&d>0?d--:40===b.keyCode&&d<c.length-1?d++:~d||(d=0);var e=c.eq(d);if(e.focus(),32===b.keyCode||13===b.keyCode){var f=e.find("input");f.prop("checked",!f.prop("checked")),f.change()}b.stopPropagation(),b.preventDefault()}},this)),this.options.enableClickableOptGroups&&this.options.multiple&&a("li.multiselect-group",this.$ul).on("click",a.proxy(function(b){b.stopPropagation();var c=a(b.target).parent(),d=c.nextUntil("li.multiselect-group"),e=d.filter(":visible:not(.disabled)"),f=!0,g=e.find("input");g.each(function(){f=f&&a(this).prop("checked")}),g.prop("checked",!f).trigger("change")},this))},createOptionValue:function(b){var c=a(b);c.is(":selected")&&c.prop("selected",!0);var d=this.options.optionLabel(b),e=c.val(),f=this.options.multiple?"checkbox":"radio",g=a(this.options.templates.li),h=a("label",g);h.addClass(f),this.options.enableHTML?h.html(" "+d):h.text(" "+d);var i=a("<input/>").attr("type",f).addClass("ace");this.options.checkboxName&&i.attr("name",this.options.checkboxName),h.prepend(i),i.after('<span class="lbl" />');var j=c.prop("selected")||!1;i.val(e),e===this.options.selectAllValue&&(g.addClass("multiselect-item multiselect-all"),i.parent().parent().addClass("multiselect-all")),h.attr("title",c.attr("title")),this.$ul.append(g),c.is(":disabled")&&i.attr("disabled","disabled").prop("disabled",!0).closest("a").attr("tabindex","-1").closest("li").addClass("disabled"),i.prop("checked",j),j&&this.options.selectedClass&&i.closest("li").addClass(this.options.selectedClass)},createDivider:function(b){var c=a(this.options.templates.divider);this.$ul.append(c)},createOptgroup:function(b){var c=a(b).prop("label"),d=a(this.options.templates.liGroup);this.options.enableHTML?a("label",d).html(c):a("label",d).text(c),this.options.enableClickableOptGroups&&d.addClass("multiselect-group-clickable"),this.$ul.append(d),a(b).is(":disabled")&&d.addClass("disabled"),a("option",b).each(a.proxy(function(a,b){this.createOptionValue(b)},this))},buildSelectAll:function(){"number"==typeof this.options.selectAllValue&&(this.options.selectAllValue=this.options.selectAllValue.toString());var b=this.hasSelectAll();if(!b&&this.options.includeSelectAllOption&&this.options.multiple&&a("option",this.$select).length>this.options.includeSelectAllIfMoreThan){this.options.includeSelectAllDivider&&this.$ul.prepend(a(this.options.templates.divider));var c=a(this.options.templates.li);a("label",c).addClass("checkbox"),this.options.enableHTML?a("label",c).html(" "+this.options.selectAllText):a("label",c).text(" "+this.options.selectAllText),this.options.selectAllName?a("label",c).prepend('<input type="checkbox" name="'+this.options.selectAllName+'" />'):a("label",c).prepend('<input type="checkbox" />');var d=a("input",c);d.val(this.options.selectAllValue),c.addClass("multiselect-item multiselect-all"),d.parent().parent().addClass("multiselect-all"),this.$ul.prepend(c),d.prop("checked",!1)}},buildFilter:function(){if(this.options.enableFiltering||this.options.enableCaseInsensitiveFiltering){var b=Math.max(this.options.enableFiltering,this.options.enableCaseInsensitiveFiltering);if(this.$select.find("option").length>=b){if(this.$filter=a(this.options.templates.filter),a("input",this.$filter).attr("placeholder",this.options.filterPlaceholder),this.options.includeFilterClearBtn){var c=a(this.options.templates.filterClearBtn);c.on("click",a.proxy(function(b){clearTimeout(this.searchTimeout),this.$filter.find(".multiselect-search").val(""),a("li",this.$ul).show().removeClass("filter-hidden"),this.updateSelectAll()},this)),this.$filter.find(".input-group").append(c)}this.$ul.prepend(this.$filter),this.$filter.val(this.query).on("click",function(a){a.stopPropagation()}).on("input keydown",a.proxy(function(b){13===b.which&&b.preventDefault(),clearTimeout(this.searchTimeout),this.searchTimeout=this.asyncFunction(a.proxy(function(){if(this.query!==b.target.value){this.query=b.target.value;var c,d;a.each(a("li",this.$ul),a.proxy(function(b,e){var f=a("input",e).length>0?a("input",e).val():"",g=a("label",e).text(),h="";if("text"===this.options.filterBehavior?h=g:"value"===this.options.filterBehavior?h=f:"both"===this.options.filterBehavior&&(h=g+"\n"+f),f!==this.options.selectAllValue&&g){var i=!1;this.options.enableCaseInsensitiveFiltering&&h.toLowerCase().indexOf(this.query.toLowerCase())>-1?i=!0:h.indexOf(this.query)>-1&&(i=!0),a(e).toggle(i).toggleClass("filter-hidden",!i),a(e).hasClass("multiselect-group")?(c=e,d=i):(i&&a(c).show().removeClass("filter-hidden"),!i&&d&&a(e).show().removeClass("filter-hidden"))}},this))}this.updateSelectAll()},this),300,this)},this))}}},destroy:function(){this.$container.remove(),this.$select.show(),this.$select.data("multiselect",null)},refresh:function(){a("option",this.$select).each(a.proxy(function(b,c){var d=a("li input",this.$ul).filter(function(){return a(this).val()===a(c).val()});a(c).is(":selected")?(d.prop("checked",!0),this.options.selectedClass&&d.closest("li").addClass(this.options.selectedClass)):(d.prop("checked",!1),this.options.selectedClass&&d.closest("li").removeClass(this.options.selectedClass)),a(c).is(":disabled")?d.attr("disabled","disabled").prop("disabled",!0).closest("li").addClass("disabled"):d.prop("disabled",!1).closest("li").removeClass("disabled")},this)),this.updateButtonText(),this.updateSelectAll()},select:function(b,c){a.isArray(b)||(b=[b]);for(var d=0;d<b.length;d++){var e=b[d];if(null!==e&&void 0!==e){var f=this.getOptionByValue(e),g=this.getInputByValue(e);void 0!==f&&void 0!==g&&(this.options.multiple||this.deselectAll(!1),this.options.selectedClass&&g.closest("li").addClass(this.options.selectedClass),g.prop("checked",!0),f.prop("selected",!0),c&&this.options.onChange(f,!0))}}this.updateButtonText(),this.updateSelectAll()},clearSelection:function(){this.deselectAll(!1),this.updateButtonText(),this.updateSelectAll()},deselect:function(b,c){a.isArray(b)||(b=[b]);for(var d=0;d<b.length;d++){var e=b[d];if(null!==e&&void 0!==e){var f=this.getOptionByValue(e),g=this.getInputByValue(e);void 0!==f&&void 0!==g&&(this.options.selectedClass&&g.closest("li").removeClass(this.options.selectedClass),g.prop("checked",!1),f.prop("selected",!1),c&&this.options.onChange(f,!1))}}this.updateButtonText(),this.updateSelectAll()},selectAll:function(b,c){var b="undefined"==typeof b?!0:b,d=a("li input[type='checkbox']:enabled",this.$ul),e=d.filter(":visible"),f=d.length,g=e.length;if(b?(e.prop("checked",!0),a("li:not(.divider):not(.disabled)",this.$ul).filter(":visible").addClass(this.options.selectedClass)):(d.prop("checked",!0),a("li:not(.divider):not(.disabled)",this.$ul).addClass(this.options.selectedClass)),f===g||b===!1)a("option:enabled",this.$select).prop("selected",!0);else{var h=e.map(function(){return a(this).val()}).get();a("option:enabled",this.$select).filter(function(b){return-1!==a.inArray(a(this).val(),h)}).prop("selected",!0)}c&&this.options.onSelectAll()},deselectAll:function(b){var b="undefined"==typeof b?!0:b;if(b){var c=a("li input[type='checkbox']:not(:disabled)",this.$ul).filter(":visible");c.prop("checked",!1);var d=c.map(function(){return a(this).val()}).get();a("option:enabled",this.$select).filter(function(b){return-1!==a.inArray(a(this).val(),d)}).prop("selected",!1),this.options.selectedClass&&a("li:not(.divider):not(.disabled)",this.$ul).filter(":visible").removeClass(this.options.selectedClass)}else a("li input[type='checkbox']:enabled",this.$ul).prop("checked",!1),a("option:enabled",this.$select).prop("selected",!1),this.options.selectedClass&&a("li:not(.divider):not(.disabled)",this.$ul).removeClass(this.options.selectedClass)},rebuild:function(){this.$ul.html(""),this.options.multiple="multiple"===this.$select.attr("multiple"),this.buildSelectAll(),this.buildDropdownOptions(),this.buildFilter(),this.updateButtonText(),this.updateSelectAll(),this.options.disableIfEmpty&&a("option",this.$select).length<=0?this.disable():this.enable(),this.options.dropRight&&this.$ul.addClass("pull-right")},dataprovider:function(c){var d=0,e=this.$select.empty();a.each(c,function(c,f){var g;a.isArray(f.children)?(d++,g=a("<optgroup/>").attr({label:f.label||"Group "+d,disabled:!!f.disabled}),b(f.children,function(b){g.append(a("<option/>").attr({value:b.value,label:b.label||b.value,title:b.title,selected:!!b.selected,disabled:!!b.disabled}))})):g=a("<option/>").attr({value:f.value,label:f.label||f.value,title:f.title,selected:!!f.selected,disabled:!!f.disabled}),e.append(g)}),this.rebuild()},enable:function(){this.$select.prop("disabled",!1),this.$button.prop("disabled",!1).removeClass("disabled")},disable:function(){this.$select.prop("disabled",!0),this.$button.prop("disabled",!0).addClass("disabled")},setOptions:function(a){this.options=this.mergeOptions(a)},mergeOptions:function(b){return a.extend(!0,{},this.defaults,this.options,b)},hasSelectAll:function(){return a("li.multiselect-all",this.$ul).length>0},updateSelectAll:function(){if(this.hasSelectAll()){var b=a("li:not(.multiselect-item):not(.filter-hidden) input:enabled",this.$ul),c=b.length,d=b.filter(":checked").length,e=a("li.multiselect-all",this.$ul),f=e.find("input");d>0&&d===c?(f.prop("checked",!0),e.addClass(this.options.selectedClass),this.options.onSelectAll()):(f.prop("checked",!1),e.removeClass(this.options.selectedClass))}},updateButtonText:function(){var b=this.getSelected();this.options.enableHTML?a(".multiselect .multiselect-selected-text",this.$container).html(this.options.buttonText(b,this.$select)):a(".multiselect .multiselect-selected-text",this.$container).text(this.options.buttonText(b,this.$select)),a(".multiselect",this.$container).attr("title",this.options.buttonTitle(b,this.$select))},getSelected:function(){return a("option",this.$select).filter(":selected")},getOptionByValue:function(b){for(var c=a("option",this.$select),d=b.toString(),e=0;e<c.length;e+=1){var f=c[e];if(f.value===d)return a(f)}},getInputByValue:function(b){for(var c=a("li input",this.$ul),d=b.toString(),e=0;e<c.length;e+=1){var f=c[e];if(f.value===d)return a(f)}},updateOriginalOptions:function(){this.originalOptions=this.$select.clone()[0].options},asyncFunction:function(a,b,c){var d=Array.prototype.slice.call(arguments,3);return setTimeout(function(){a.apply(c||window,d)},b)},setAllSelectedText:function(a){this.options.allSelectedText=a,this.updateButtonText()}},a.fn.multiselect=function(b,d,e){return this.each(function(){var f=a(this).data("multiselect"),g="object"==typeof b&&b;f||(f=new c(this,g),a(this).data("multiselect",f)),"string"==typeof b&&(f[b](d,e),"destroy"===b&&a(this).data("multiselect",!1))})},a.fn.multiselect.Constructor=c,a(function(){a("select[data-role=multiselect]").multiselect()})}(window.jQuery);
!function(a){"function"==typeof define&&define.amd?define(["jquery"],a):"object"==typeof exports?module.exports=a(require("jquery")):a(jQuery)}(function(a){var b=a.fn.spinbox,c=function(b,c){this.$element=a(b),this.$element.find(".btn").on("click",function(a){a.preventDefault()}),this.options=a.extend({},a.fn.spinbox.defaults,c),this.options.step=this.$element.data("step")||this.options.step,this.options.value<this.options.min?this.options.value=this.options.min:this.options.max<this.options.value&&(this.options.value=this.options.max),this.$input=this.$element.find(".spinbox-input"),this.$input.on("focusout.fu.spinbox",this.$input,a.proxy(this.change,this)),this.$element.on("keydown.fu.spinbox",this.$input,a.proxy(this.keydown,this)),this.$element.on("keyup.fu.spinbox",this.$input,a.proxy(this.keyup,this)),this.bindMousewheelListeners(),this.mousewheelTimeout={},this.options.hold?(this.$element.on("mousedown.fu.spinbox",".spinbox-up",a.proxy(function(){this.startSpin(!0)},this)),this.$element.on("mouseup.fu.spinbox",".spinbox-up, .spinbox-down",a.proxy(this.stopSpin,this)),this.$element.on("mouseout.fu.spinbox",".spinbox-up, .spinbox-down",a.proxy(this.stopSpin,this)),this.$element.on("mousedown.fu.spinbox",".spinbox-down",a.proxy(function(){this.startSpin(!1)},this))):(this.$element.on("click.fu.spinbox",".spinbox-up",a.proxy(function(){this.step(!0)},this)),this.$element.on("click.fu.spinbox",".spinbox-down",a.proxy(function(){this.step(!1)},this))),this.switches={count:1,enabled:!0},"medium"===this.options.speed?this.switches.speed=300:"fast"===this.options.speed?this.switches.speed=100:this.switches.speed=500,this.options.defaultUnit=e(this.options.defaultUnit,this.options.units)?this.options.defaultUnit:"",this.unit=this.options.defaultUnit,this.lastValue=this.options.value,this.render(),this.options.disabled&&this.disable()},d=function(a,b){return Math.round(a/b)*b},e=function(b,c){var d=!1,e=b.toLowerCase();return a.each(c,function(a,b){return b=b.toLowerCase(),e===b?(d=!0,!1):void 0}),d},f=function(a){return isNaN(parseFloat(a))?a:(a>this.options.max?a=this.options.cycle?this.options.min:this.options.max:a<this.options.min&&(a=this.options.cycle?this.options.max:this.options.min),this.options.limitToStep&&this.options.step&&(a=d(a,this.options.step),a>this.options.max?a-=this.options.step:a<this.options.min&&(a+=this.options.step)),a)};c.prototype={constructor:c,destroy:function(){return this.$element.remove(),this.$element.find("input").each(function(){a(this).attr("value",a(this).val())}),this.$element[0].outerHTML},render:function(){this.setValue(this.getDisplayValue())},change:function(){this.setValue(this.getDisplayValue()),this.triggerChangedEvent()},stopSpin:function(){void 0!==this.switches.timeout&&(clearTimeout(this.switches.timeout),this.switches.count=1,this.triggerChangedEvent())},triggerChangedEvent:function(){var a=this.getValue();a!==this.lastValue&&(this.lastValue=a,this.$element.trigger("changed.fu.spinbox",a))},startSpin:function(b){if(!this.options.disabled){var c=this.switches.count;1===c?(this.step(b),c=1):c=3>c?1.5:8>c?2.5:4,this.switches.timeout=setTimeout(a.proxy(function(){this.iterate(b)},this),this.switches.speed/c),this.switches.count++}},iterate:function(a){this.step(a),this.startSpin(a)},step:function(a){this.setValue(this.getDisplayValue());var b;b=a?this.options.value+this.options.step:this.options.value-this.options.step,b=b.toFixed(5),this.setValue(b+this.unit)},getDisplayValue:function(){var a=this.parseInput(this.$input.val()),b=a?a:this.options.value;return b},setDisplayValue:function(a){this.$input.val(a)},getValue:function(){var a=this.options.value;return"."!==this.options.decimalMark&&(a=(a+"").split(".").join(this.options.decimalMark)),a+this.unit},setValue:function(a){if("."!==this.options.decimalMark&&(a=this.parseInput(a)),"number"!=typeof a){var b=a.replace(/[0-9.-]/g,"");this.unit=e(b,this.options.units)?b:this.options.defaultUnit}var c=this.getIntValue(a);return isNaN(c)&&!isFinite(c)?this.setValue(this.options.value):(c=f.call(this,c),this.options.value=c,a=c+this.unit,"."!==this.options.decimalMark&&(a=(a+"").split(".").join(this.options.decimalMark)),this.setDisplayValue(a),this)},value:function(a){return a||0===a?this.setValue(a):this.getValue()},parseInput:function(a){return a=(a+"").split(this.options.decimalMark).join(".")},getIntValue:function(a){return a="undefined"==typeof a?this.getValue():a,"undefined"!=typeof a?("string"==typeof a&&(a=this.parseInput(a)),a=parseFloat(a,10)):void 0},disable:function(){this.options.disabled=!0,this.$element.addClass("disabled"),this.$input.attr("disabled",""),this.$element.find("button").addClass("disabled")},enable:function(){this.options.disabled=!1,this.$element.removeClass("disabled"),this.$input.removeAttr("disabled"),this.$element.find("button").removeClass("disabled")},keydown:function(a){var b=a.keyCode;38===b?this.step(!0):40===b?this.step(!1):13===b&&this.change()},keyup:function(a){var b=a.keyCode;(38===b||40===b)&&this.triggerChangedEvent()},bindMousewheelListeners:function(){var b=this.$input.get(0);b.addEventListener?(b.addEventListener("mousewheel",a.proxy(this.mousewheelHandler,this),!1),b.addEventListener("DOMMouseScroll",a.proxy(this.mousewheelHandler,this),!1)):b.attachEvent("onmousewheel",a.proxy(this.mousewheelHandler,this))},mousewheelHandler:function(a){if(!this.options.disabled){var b=window.event||a,c=Math.max(-1,Math.min(1,b.wheelDelta||-b.detail)),d=this;return clearTimeout(this.mousewheelTimeout),this.mousewheelTimeout=setTimeout(function(){d.triggerChangedEvent()},300),c>0?this.step(!0):this.step(!1),b.preventDefault?b.preventDefault():b.returnValue=!1,!1}}},a.fn.spinbox=function(b){var d,e=Array.prototype.slice.call(arguments,1),f=this.each(function(){var f=a(this),g=f.data("fu.spinbox"),h="object"==typeof b&&b;g||f.data("fu.spinbox",g=new c(this,h)),"string"==typeof b&&(d=g[b].apply(g,e))});return void 0===d?f:d},a.fn.spinbox.defaults={value:0,min:0,max:999,step:1,hold:!0,speed:"medium",disabled:!1,cycle:!1,units:[],decimalMark:".",defaultUnit:"",limitToStep:!1},a.fn.spinbox.Constructor=c,a.fn.spinbox.noConflict=function(){return a.fn.spinbox=b,this},a(document).on("mousedown.fu.spinbox.data-api","[data-initialize=spinbox]",function(b){var c=a(b.target).closest(".spinbox");c.data("fu.spinbox")||c.spinbox(c.data())}),a(function(){a("[data-initialize=spinbox]").each(function(){var b=a(this);b.data("fu.spinbox")||b.spinbox(b.data())})})});