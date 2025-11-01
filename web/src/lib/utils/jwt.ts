export function getUserIdFromToken(token: string): number | null {
	try {
		const payload = JSON.parse(atob(token.split('.')[1]));
		return payload.userID;
	} catch {
		return null;
	}
}

export function getUserRoleFromToken(token: string): string | null {
	try {
		const payload = JSON.parse(atob(token.split('.')[1]));
		if (payload.role == 1) {
			return 'user';
		} else if (payload.role == 2) {
			return 'moderator';
		} else if (payload.role == 3) {
			return 'admin';
		} else {
			return null;
		}
	} catch {
		return null;
	}
}
