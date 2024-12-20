import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	// const res = await fetch(`http://localhost:8080/api/item/${params.item}`);
	// const data = await res.json();
	// console.log(data);

	return {
		item: params.item,
        gender: params.gender
	};
};
