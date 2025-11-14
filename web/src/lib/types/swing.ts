export type Swing = {
	ID: number;
	UserID: number;
	Name: string;
	Description: string;
	IsAccessible: boolean;
	City: string;
	Address: string;
	Lat?: number;
	Lng?: number;
	SeatCount: number;
	MaxWeightKg: number;
	CreatedAt: string;
	UpdatedAt: string;
	Tags: Tag[];
	Materials: Material[];
};

type Tag = {
	Name: string;
};

type Material = {
	ID: number;
	Name: string;
};
