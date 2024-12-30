import { Button } from "@/components/ui/button";
import type React from "react";

export const About: React.FC = () => {
	return (
		<div className="min-h-screen w-screen flex justify-center items-center">
			<div className="flex flex-col justify-center items-center">
				<h1 className="text-3xl font-bold underline">About Page</h1>
				<Button className="mt-3">Click me</Button>
			</div>
		</div>
	);
};
