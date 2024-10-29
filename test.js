import { Generate } from 'k6/x/totp'; 

export default function () {
    const secret = 'JBSWY3DPEHPK3PXP'; // Example TOTP secret (Base32)
    const token = Generate(secret); // Ensure this matches your method name
    console.log(`Generated TOTP: ${token}`); // Log the generated token
}
