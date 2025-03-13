/**
 * Landing Page - Features Section
 */

// Dependencies
import { motion } from 'framer-motion';

const FEATURES_CONTENT = [
	{
		title: 'Free Test and Connect',
		description:
			'Take advantage of our free and comprehensive test. Answer the questions, and connect to similar people on your emotional well-being without any underlying cost.',
		available: true,
		image: '/images/features/test.svg',
	},
	{
		title: 'Stay Anonymous',
		description:
			'Protect your privacy and stay anonymous. Share your thoughts and feelings without revealing your identity. Connect with others who understand your emotional state and experiences.',
		available: true,
		image: '/images/features/history.svg',
	},
	{
		title: 'Login/Create Account to Store Test Data',
		description:
			'Create your account or log in to store your test data securely. Your information will be protected, allowing you to track your progress over time and access your past test results whenever you need them. Securely store your test data, export it for personal records, and exercise control over your mental health journey. If needed, seamlessly delete your account.',
		available: true,
		image: '/images/features/account.svg',
	},
	{
		title: "Relatable Articles and Blogs",
		description:
			'A treasure trove of articles and blogs, personalized to your mental health test results. Unlock valuable insights, strategies, and self-care tips, empowering you to thrive on your emotional well-being journey.',
		available: false,
		image: '/images/features/resources.svg',
	},
];

function Features() {
	return (
		<motion.section
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ delay: 0.2, type: 'spring' }}
			viewport={{ once: true }}
			className='mt-20 text-center md:text-left scroll-m-20'
			id='features'
		>
			<h2 className='font-heading text-2xl md:text-4xl font-bold'>
				Our Features
			</h2>
			<ul className='grid grid-cols-1 md:grid-cols-2 gap-8 items-center mt-4'>
				{FEATURES_CONTENT.map((content, index) => (
					<motion.li
						initial={{ opacity: 0, scale: 1.1, y: -20 }}
						whileInView={{ opacity: 1, scale: 1, y: 0 }}
						transition={{
							delay: 0.1 * (index + 1),
							type: 'keyframes',
							duration: 0.2,
						}}
						viewport={{ once: true }}
						key={`${content.title}-${index}`}
						className='border-2 border-secondary overflow-hidden group transition-all hover:border-secondaryDark rounded-3xl px-8 py-4 h-full select-none hover:-translate-y-2 hover:-translate-x-2'
					>
						<h3 className='font-heading text-lg md:text-2xl font-semibold'>
							{content.title}
						</h3>
						<p>{content.available ? 'Ready' : 'Coming Soon'}</p>
						<div className='w-[160px] h-[160px] p-4 flex items-center justify-center mx-auto relative my-4'>
							<img
								src={content.image}
								alt={content.title}
								className='w-full h-auto object-contain z-10 group-hover:scale-95 transition-all duration-700'
								loading='lazy'
							/>
							<div className='bg-secondary rounded-full w-[160px] h-[160px] group-hover:w-[500%] group-hover:h-[500%] bg-opacity-40 absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2 z-0 transition-all duration-300' />
						</div>
						<p>{content.description}</p>
					</motion.li>
				))}
			</ul>
		</motion.section>
	);
}

export default Features;
