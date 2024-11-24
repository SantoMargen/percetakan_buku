const CryptoJS = require('crypto-js');


const iv = CryptoJS.lib.WordArray.random(16);
const key = CryptoJS.enc.Utf8.parse('1n1S4ng4tr4has14');

// const plaintext = {"request":{"full_name":"Bunga Cantika Larasati","email":"bungacantika123@gmail.com","password":"bunga123","phone_number":"082299024","role":"EDITOR 2","gender":"M"}}
// const plaintext = {"request":{"page":1,"size":5,"filter":{"title":"","publish_date":""}}}
const plaintext = {
  "request":{
    "email": "admin3@mail.com",
    "password":"P@ssw0rd",
}
}

// const plaintext = {"request":{"ID":65}}

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
console.log("DECRYPT: ", Decrypt("n1f3a8hYKBTmC8qU0s0ZUWzSSULyk/vRp/bY5tgVFsQFUi9yZu5NazdtPo5PVK8x"))