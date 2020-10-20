(window.webpackJsonp=window.webpackJsonp||[]).push([[27],{82:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return a})),n.d(t,"metadata",(function(){return c})),n.d(t,"rightToc",(function(){return l})),n.d(t,"default",(function(){return p}));var r=n(2),o=n(6),i=(n(0),n(90)),a={id:"installation.contributing",title:"Contributing",sidebar_label:"Contributing",slug:"/installation/contributing"},c={unversionedId:"installation.contributing",id:"installation.contributing",isDocsHomePage:!1,title:"Contributing",description:"If you wish to work on Semaphore itself or any of its built-in systems, you'll",source:"@site/docs/installation-contributing.md",slug:"/installation/contributing",permalink:"/semaphore/docs/installation/contributing",editUrl:"https://github.com/jexia/semaphore/edit/master/website/docs/installation-contributing.md",version:"current",sidebar_label:"Contributing",sidebar:"docs",previous:{title:"Build from source",permalink:"/semaphore/docs/installation/source"},next:{title:"Getting started",permalink:"/semaphore/docs/flows"}},l=[],s={rightToc:l};function p(e){var t=e.components,n=Object(o.a)(e,["components"]);return Object(i.b)("wrapper",Object(r.a)({},s,n,{components:t,mdxType:"MDXLayout"}),Object(i.b)("p",null,"If you wish to work on Semaphore itself or any of its built-in systems, you'll\nfirst need ",Object(i.b)("a",Object(r.a)({parentName:"p"},{href:"https://www.golang.org"}),"Go")," installed on your machine. Go version\n1.13.7+ is ",Object(i.b)("em",{parentName:"p"},"required"),"."),Object(i.b)("p",null,"For local dev first make sure Go is properly installed, including setting up a\n",Object(i.b)("a",Object(r.a)({parentName:"p"},{href:"https://golang.org/doc/code.html#GOPATH"}),"GOPATH"),". Ensure that ",Object(i.b)("inlineCode",{parentName:"p"},"$GOPATH/bin")," is in\nyour path as some distributions bundle old version of build tools. Next, clone this\nrepository. Semaphore uses ",Object(i.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/golang/go/wiki/Modules"}),"Go Modules"),",\nso it is recommended that you clone the repository ",Object(i.b)("strong",{parentName:"p"},Object(i.b)("em",{parentName:"strong"},"outside"))," of the GOPATH.\nYou can then download any required build tools by bootstrapping your environment:"),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{className:"language-sh"}),"$ make bootstrap\n...\n")),Object(i.b)("p",null,"To compile a development version of Semaphore, run ",Object(i.b)("inlineCode",{parentName:"p"},"make")," or ",Object(i.b)("inlineCode",{parentName:"p"},"make dev"),". This will\nput the Semaphore binary in the ",Object(i.b)("inlineCode",{parentName:"p"},"bin")," folders:"),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{className:"language-sh"}),"$ make dev\n...\n$ bin/semaphore\n...\n")),Object(i.b)("p",null,"To run tests, type ",Object(i.b)("inlineCode",{parentName:"p"},"make test"),". If\nthis exits with exit status 0, then everything is working!"),Object(i.b)("pre",null,Object(i.b)("code",Object(r.a)({parentName:"pre"},{className:"language-sh"}),"$ make test\n...\n")))}p.isMDXComponent=!0},90:function(e,t,n){"use strict";n.d(t,"a",(function(){return u})),n.d(t,"b",(function(){return d}));var r=n(0),o=n.n(r);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=o.a.createContext({}),p=function(e){var t=o.a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):c(c({},t),e)),n},u=function(e){var t=p(e.components);return o.a.createElement(s.Provider,{value:t},e.children)},b={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},m=o.a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,a=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=p(n),m=r,d=u["".concat(a,".").concat(m)]||u[m]||b[m]||i;return n?o.a.createElement(d,c(c({ref:t},s),{},{components:n})):o.a.createElement(d,c({ref:t},s))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,a=new Array(i);a[0]=m;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c.mdxType="string"==typeof e?e:r,a[1]=c;for(var s=2;s<i;s++)a[s]=n[s];return o.a.createElement.apply(null,a)}return o.a.createElement.apply(null,n)}m.displayName="MDXCreateElement"}}]);