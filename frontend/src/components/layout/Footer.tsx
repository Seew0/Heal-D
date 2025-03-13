/**
 * Footer Component
 */

// Dependencies
import { Link } from 'react-router-dom';
import { FOOTER_LINKS } from '../../data/navigation';

function Footer() {
	return (
		<footer className='py-4 px-8'>
			<div className='max-w-7xl mx-auto flex items-center gap-8 justify-between flex-wrap'>
				<Link
					to={'/'}
					className='flex w-fit items-center gap-1 hover:scale-[0.98] active:scale-[1.02] transition-all'
				>
					<div className='w-10 h-10'>
						<img
							src='/mind-check-logo.png'
							alt='🧠'
							className='w-full h-auto object-contain'
							loading='lazy'
						/>
					</div>
					<p className='font-heading text-3xl font-bold'>
						Heal'D
					</p>
				</Link>
				<ul className='grid grid-cols-3 gap-4'>
					{FOOTER_LINKS.map((link, index) => (
						<li
							key={index}
							className='text-textSecondary transition-all font-semibold text-lg hover:underline'
						>
							<Link to={link.url}>{link.name}</Link>
						</li>
					))}
				</ul>
				<hr className='w-full border' />
				<p className='text-center w-full'>
					Copyright &copy; Heal'D | Made with ❤️
					| {" "}
					<a
						href='https://github.com/seew0/Heal-D'
						target='_blank'
						className='text-textSecondary underline text-opacity-80 hover:text-opacity-100 transition-all'
					>
						GitHub
					</a>{' '}
					| {" "}
					<a>
						Created By Devansh Miglani
					</a>
				</p>
			</div>
		</footer>
	);
}

export default Footer;
