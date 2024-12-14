import type { PageLoad } from './$types';

export const load: PageLoad = async ({fetch, params }) => {
    const res = await fetch(`http://localhost:8080/api/clothes/${params.gender}`);
    const data = await res.json();
    console.log(data);

	return {
		slug: params.gender,
        items: data
	};
};
