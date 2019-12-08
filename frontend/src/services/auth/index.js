import auth0 from 'auth0-js';
import { DOMAIN, CLIENT_ID } from './config';

const webAuth = new auth0.WebAuth({
  domain: DOMAIN,
  redirectUri: `${window.location.origin}/callback`,
  clientID: CLIENT_ID,
  audience: `https://${DOMAIN}/api/v2/`,
  responseType: 'id_token token',
  scope: 'openid profile email app_metadata',
});

export const handleAuthentication = () => new Promise((resolve, reject) => {
  webAuth.parseHash((err, authResult) => {
    if (err) {
      reject(err);
    } else {
      resolve({
        idToken: authResult.idToken,
        name: authResult.idTokenPayload.name,
        tokenExpiry: authResult.idTokenPayload.exp,
      });
    }
  });
});

export const authorize = () => webAuth.authorize();
