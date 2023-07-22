htmx.defineExtension("renderProp", {
  onEvent: function (info, e) {
    if (info === "htmx:afterRequest") {
      var target = e.detail.target;

      var res = JSON.parse(e.detail.xhr.response);
      console.log(res);
      const arrPos = +e.target.attributes.propName.value.charAt(
        e.target.attributes.propName.value.indexOf("[") + 1
      );

      let prop;

      const splited = e.target.attributes.propName.value.split(".");

      if (!isNaN(arrPos)) {
        const firstProp = e.target.attributes.propName.value.split("[")[0];
        prop = res[firstProp][arrPos];
        splited.shift();
      } else {
        prop = res;
      }

      for (let i = 0; i < splited.length; i++) {
        prop = prop[splited[i]];
      }

      target.innerHTML = prop;
      console.log(prop);
    }
  },
});
