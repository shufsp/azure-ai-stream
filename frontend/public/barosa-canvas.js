function getResultValuesFromConfidence(json, confidence) {
   if (!json) {
      return []
   }

   const az = json.azureResponse
   const resultJson = (az.denseCaptionsResult || az.peopleResult).values
   if (!resultJson) {
      console.warn(`Result json doesn't have bounding box objects (azure response: ${json})`)
      return []
   }

   const resultValues = resultJson
      .filter(k => k["confidence"] >= confidence)
   return resultValues
}

function drawBoundingBox(resultValue, ctx) {
   const { x, y, w, h } = resultValue.boundingBox;

   ctx.beginPath();
   ctx.rect(x, y, w, h);
   ctx.lineWidth = 3;
   ctx.strokeStyle = 'red';
   ctx.stroke();

   ctx.font = '12px Arial';
   ctx.fillStyle = 'red';
   ctx.fillText(resultValue.text || "Person", x, y > 10 ? y - 5 : y + 15);
}

function consumeScarabToCanvas(response, ctx) {
   const bbs = getResultValuesFromConfidence(response, 0.60)
   bbs.forEach(resultValue => drawBoundingBox(resultValue, ctx))
}

