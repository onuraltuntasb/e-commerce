// place files you want to import through the `$lib` alias in this folder.
import { z } from 'zod';

export const signupSchema = z
	.object({
		firstName: z.string().min(1),
		lastName: z.string().min(1),
		email: z.string().email(),
		//TODO password validation must be detailed
		password: z.string().min(8),
		passwordConfirm: z.string().min(8)
	})
	.refine((data) => data.password == data.passwordConfirm, {
		message: "Passwords didn't match",
		path: ['passwordConfirm']
	});

export const loginSchema = z.object({
	email: z.string().email(),
	//TODO password validation must be detailed
	password: z.string().min(8)
});


export const addressSchema = z.object({
	header: z.string().trim().min(1, { message: "Required" }),
	name: z.string().trim().min(1, { message: "Required" }),
	country: z.string().trim().min(1, { message: "Required" }),
	city: z.string().trim().min(1, { message: "Required" }),
	district: z.string().trim().min(1, { message: "Required" }),
	neighborhood: z.string().trim().min(1, { message: "Required" }),
	description: z.string().trim().min(1, { message: "Required" }),
	postalCode: z.string().trim().min(1, { message: "Required" }),
	phone: z.string().min(10),
	billType: z.boolean(),
	ssn: z.string().min(11)
});


