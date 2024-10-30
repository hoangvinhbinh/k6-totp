import totp from 'k6/x/xk6_totp';

export default function () {
  console.log(totp.generate("Daniel"));
}