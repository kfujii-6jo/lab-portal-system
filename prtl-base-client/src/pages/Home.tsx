import { Button } from "@/components/ui/button";
import React, { useState, useEffect } from "react";
import { Layout } from "@/components/layout";

type Post = {
	id: number;
	title: string;
};

export const Home: React.FC = () => {
	const [posts, setPosts] = useState<Post[]>([]);

	useEffect(() => {
		fetch("https://jsonplaceholder.typicode.com/posts", { method: "GET" })
			.then((res) => res.json())
			.then((data) => {
				setPosts(data);
			});
	}, []);
	return (
		<Layout>
			<div>
				<h1 className="text-3xl font-bold underline">Home</h1>
				<Button className="mt-3">Click me</Button>
				<ul>
					{" "}
					{posts.map((post) => (
						<li key={post.id}>
							{post.id}
							{post.title}
						</li>
					))}{" "}
				</ul>
			</div>
		</Layout>
	);
};
