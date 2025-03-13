import { createContext, ReactNode, useEffect, useState } from 'react';
import { useUserStore } from '../store/user';
import { useNavigate } from 'react-router-dom';
import { fetchAllResources } from '../utils/resources';
import LoadingScreen from '../components/reusable/LoadingScreen';

export const AuthContext = createContext({});

export const AuthProvider = ({ children }: { children: ReactNode }) => {
	const [initialLoader, setInitialLoader] = useState<boolean>(true);
	const { setUser } = useUserStore();
	const navigate = useNavigate();

	useEffect(
		() => {
			const user = localStorage.getItem('user');
			if (user) {
				setUser(JSON.parse(user));
			}
			setInitialLoader(false);
		},
		// eslint-disable-next-line
		[]

	);

	useEffect(() => {
		fetchAllResources({});
	}, []);

	return (
		<AuthContext.Provider value={{}}>
			{initialLoader ? <LoadingScreen /> : children}
		</AuthContext.Provider>
	);
};
