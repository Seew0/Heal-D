/**
 * Navigation Data
 */

interface NavLink {
	name: string;
	url: string;
}

type NavLinks = NavLink[];

const COMMON_LINKS: NavLinks = [
	{ name: 'Home', url: '/' },
	{ name: 'Blogs', url: '/resources' },
];

export const MAIN_LINKS: NavLinks = [...COMMON_LINKS];

export const FOOTER_LINKS: NavLinks = [
	...COMMON_LINKS
];
