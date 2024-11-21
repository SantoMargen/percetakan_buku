const CryptoJS = require('crypto-js');


const iv = CryptoJS.lib.WordArray.random(16);
const key = CryptoJS.enc.Utf8.parse('1n1S4ng4tr4has14');

// const plaintext = {"request":{"full_name":"Bunga Cantika Larasati","email":"bungacantika123@gmail.com","password":"bunga123","phone_number":"082299024","role":"EDITOR 2","gender":"M"}}
const plaintext = {"request":{"page":1,"size":5,"filter":{"user":"18309","title":"","publish_date":""}}}
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
// console.log("DECRYPT: ", Decrypt("igc+JVXqsw/hz8dlDKKITcKe+dOnhOnVP6+8d01y7TX4QeH+bYAai1fu0vjRk5QffzN7hnBCr+0SXGc8psT/DmYCc3vg+bPWvPnTNeSyp+jhW2IFEPP2tcwjsL2z342t"))