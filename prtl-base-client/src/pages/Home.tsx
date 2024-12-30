import { Button } from "@/components/ui/button";
import type React from "react";
import { Layout } from "@/components/layout";

export const Home: React.FC = () => {
	return (
		<Layout>
			<div>
				<h1 className="text-3xl font-bold underline">Home</h1>
				<Button className="mt-3">Click me</Button>
			</div>
		</Layout>
	);
};
