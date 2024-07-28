
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

function flattenObjectToUrlQuery(obj, prefix = '') {
  const queryString = [];

  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      const newPrefix = prefix ? `${prefix}${key}` : `${key}`;
      if (typeof obj[key] === 'object' && obj[key] !== null && !Array.isArray(obj[key])) {
        queryString.push(flattenObjectToUrlQuery(obj[key], newPrefix));
      } else {
        queryString.push(`${newPrefix}=${encodeURIComponent(obj[key])}`);
      }
    }
  }

  return queryString.join('&');
}

function createImageFeaturesUrl(props) {
  const flatcake = flattenObjectToUrlQuery(props)
  return `image-features?${flatcake}`
}

async function scarab(props) {
  return new Promise(async (resolve, _) => {
    const token = props.token;
    delete props.token;

    const pingResponse = await requestFromBarosaServer("ping", token)
    if (pingResponse.error) {
      console.error("Couldn't reach the barosa server.. is the GO server started?")
      return
    }

    const imgFeaturesRequestUrl = createImageFeaturesUrl(props)
    const imgFeaturesResponse = await requestFromBarosaServer(imgFeaturesRequestUrl, token)
    resolve(imgFeaturesResponse)
  })
}
