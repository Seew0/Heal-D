/**
 * Landing Page - Intro Section
 */

// Dependencies
import { motion } from 'framer-motion';
import Lottie from 'lottie-react';
import { Link } from 'react-router-dom';
import MentalWellBeingAnimationData from '../../assets/lottie/mental-wellbeing-seek-help.json';

function Intro() {
	return (
		<motion.section
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ delay: 0.2, type: 'spring' }}
			viewport={{ once: true }}
			className='w-full grid grid-cols-1 md:grid-cols-2 place-items-center text-center md:text-left'
		>
			<div className='flex flex-col gap-4'>
				<h1 className='font-heading text-3xl md:text-5xl font-bold'>
				Join a community that understands you. Talk, Heal, and Grow—Together.
				</h1>
				<h2 className='text-lg md:text-2xl font-medium max-w-2xl'>
					Take a simple test to understand your mental
					state. Heal'D is an anonymous mental health support platform that connects people based on their mental state and experiences. It provides a safe space where users can join forums with others who share similar emotional
				</h2>
				<Link
					to='/test'
					className='px-8 py-4 w-fit mx-auto md:mx-0 border-secondary border-2 rounded-full font-semibold hover:bg-tertiary transition-all hover:border-secondaryDark'
				>
					Start Your Assessment
				</Link>
			</div>
			<div className='flex items-center justify-center'>
				<div className='max-w-md md:max-w-lg'>
					<Lottie
						animationData={MentalWellBeingAnimationData}
						loop={false}
						className='w-full h-auto object-contain'
					/>
				</div>
			</div>
		</motion.section>
	);
}

export default Intro;
