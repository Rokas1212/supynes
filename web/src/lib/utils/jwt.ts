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

export function getUserRoleFromRoleNumber(role: string): string | null {
	if (role == '1') {
		return 'User';
	} else if (role == '2') {
		return 'Moderator';
	} else if (role == '3') {
		return 'Admin';
	} else {
		return null;
	}
}
