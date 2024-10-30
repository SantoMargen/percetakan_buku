const CryptoJS = require('crypto-js');


const iv = CryptoJS.lib.WordArray.random(16);
const key = CryptoJS.enc.Utf8.parse('1n1S4ng4tr4has14');

const plaintext = {"request":{"ID":2}}
// const plaintext = {"request":
//   {"name": "Publisher 1", "address": "123 Main St, City A, Country A", "phone": "123-456-7890", "email": "publisher1@example.com", "website": "https://publisher1.com", "founded_year": 2000, "country": "Country A", "contact_person_1": "Alice Smith", "contact_person_2": "John Doe", "fax": "123-456-7890", "social_fb_links": "https://facebook.com/publisher1", "social_twitter_links": "https://twitter.com/publisher1", "social_web_links": "https://publisher1.com/social", "join_date": "2023-01-01"}
// }
// const plaintext = {"request":
//     {"category_name": "Ilussion", "description": "123 Main St, City A, Country A", "entry_user": "Hisbikal","ID":22}
//   }


function Encrypt(plaintext) {
  const strText = JSON.stringify(plaintext);

  const encrypted = CryptoJS.AES.encrypt(strText, key, {
    iv: iv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });

  const ciphertext = iv.concat(encrypted.ciphertext).toString(CryptoJS.enc.Base64);
  return ciphertext
}

function Decrypt(ciphertext) {
  const ciphertextDecoded = CryptoJS.enc.Base64.parse(ciphertext);
  const iv = ciphertextDecoded.clone();
  iv.sigBytes = 16;
  iv.clamp();

  const encryptedMessage = ciphertextDecoded.clone();
  encryptedMessage.words.splice(0, 4);
  encryptedMessage.sigBytes -= 16;

  const decrypted = CryptoJS.AES.decrypt(
    {
      ciphertext: encryptedMessage,
    },
    key,
    {
      iv: iv,
      mode: CryptoJS.mode.CBC,
      padding: CryptoJS.pad.Pkcs7,
    }
  );

  const plaintext = decrypted.toString(CryptoJS.enc.Utf8);
  return plaintext;
}



console.log("START =======================");

console.log("ENCRYPT: ", Encrypt(plaintext));
// console.log("DECRYPT: ", Decrypt("StHMXv9196HvBzkHGI+sI0CgTC5D3TrE//SA12bGWiO2jsxHpmpffIIP/WIxZYua"))