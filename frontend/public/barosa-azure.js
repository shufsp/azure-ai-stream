
const BASE_URL = `http://localhost:8080`

function requestFromBarosaServer(request, token) {
  return new Promise(async (resolve, reject) => {
    fetch(`${BASE_URL}/${request}`, {
      headers: {
        "Authorization": `Bearer ${token}`
      }
    })
      .then(response => resolve(response))
      .catch(err => {
        console.error(`Error requesting from barosa server ${err}`)
        reject(err)
      })
  })
}

async function scarab(token, window, method, props) {
  return new Promise(async (resolve, _) => {
    const pingResponse = await requestFromBarosaServer("ping", token)
    if (pingResponse.error) {
      console.error("Couldn't reach the barosa server.. is the GO server started?")
      return
    }

    const imgFeaturesRequestUrl = `image-features?window=${window}&method=${method || "class"}&avifQuality=${props.avifQuality || ""}&avifAlphaQuality=${props.avifAlphaQuality || ""}&avifSpeed=${props.avifSpeed || ""}&features=${props.features || ""}&lanzcosWidth=${props.lanzcosWidth}`
    const imgFeaturesResponse = await requestFromBarosaServer(imgFeaturesRequestUrl, token)
    resolve(imgFeaturesResponse)
  })
}
