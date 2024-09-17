import { LucideProps } from "lucide-react";

export const Icons = {
	arrInCircle: (props: LucideProps) => (
		<svg
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<g id="icon/arrow-down-circle">
				<path
					id="Union"
					fillRule="evenodd"
					clipRule="evenodd"
					d="M12 2.5C6.75329 2.5 2.5 6.75329 2.5 12C2.5 17.2467 6.75329 21.5 12 21.5C17.2467 21.5 21.5 17.2467 21.5 12C21.5 6.75329 17.2467 2.5 12 2.5ZM1.5 12C1.5 6.20101 6.20101 1.5 12 1.5C17.799 1.5 22.5 6.20101 22.5 12C22.5 17.799 17.799 22.5 12 22.5C6.20101 22.5 1.5 17.799 1.5 12ZM12 7.5C12.2761 7.5 12.5 7.72386 12.5 8V14.7929L15.6464 11.6464C15.8417 11.4512 16.1583 11.4512 16.3536 11.6464C16.5488 11.8417 16.5488 12.1583 16.3536 12.3536L12.3536 16.3536C12.1583 16.5488 11.8417 16.5488 11.6464 16.3536L7.64645 12.3536C7.45118 12.1583 7.45118 11.8417 7.64645 11.6464C7.84171 11.4512 8.15829 11.4512 8.35355 11.6464L11.5 14.7929V8C11.5 7.72386 11.7239 7.5 12 7.5Z"
					fill="white"
				/>
			</g>
		</svg>
	),
	chevronDown: (props: LucideProps) => (
		<svg
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<g id="icon/chevron-down">
				<path
					id="Union"
					fillRule="evenodd"
					clipRule="evenodd"
					d="M5.64645 8.64651C5.84171 8.45125 6.15829 8.45125 6.35355 8.64651L12 14.293L17.6464 8.64651C17.8417 8.45125 18.1583 8.45125 18.3536 8.64651C18.5488 8.84177 18.5488 9.15835 18.3536 9.35361L12.3536 15.3536C12.1583 15.5489 11.8417 15.5489 11.6464 15.3536L5.64645 9.35361C5.45118 9.15835 5.45118 8.84177 5.64645 8.64651Z"
					fill="#E5EEFF"
					fillOpacity="0.6"
				/>
			</g>
		</svg>
	),
	arrLeft: (props: LucideProps) => (
		<svg
			width="32"
			height="32"
			viewBox="0 0 32 32"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<path
				fillRule="evenodd"
				clipRule="evenodd"
				d="M16.3536 8.64645C16.5488 8.84171 16.5488 9.15829 16.3536 9.35355L10.2071 15.5H23C23.2761 15.5 23.5 15.7239 23.5 16C23.5 16.2761 23.2761 16.5 23 16.5H10.2071L16.3536 22.6464C16.5488 22.8417 16.5488 23.1583 16.3536 23.3536C16.1583 23.5488 15.8417 23.5488 15.6464 23.3536L8.64645 16.3536C8.45118 16.1583 8.45118 15.8417 8.64645 15.6464L15.6464 8.64645C15.8417 8.45118 16.1583 8.45118 16.3536 8.64645Z"
				fill="white"
			/>
		</svg>
	),
	cosmos: (props: LucideProps) => (
		<svg
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<circle cx="12" cy="12" r="12" fill="black" />
			<path
				fillRule="evenodd"
				clipRule="evenodd"
				d="M13.4626 9.5356L9.58638 13.412C9.53639 13.315 9.49129 13.2139 9.45085 13.1091C9.31548 12.7569 9.2478 12.3832 9.2478 11.9886C9.2478 11.5936 9.31548 11.22 9.45085 10.8678C9.58638 10.5155 9.77409 10.2098 10.0141 9.95031C10.2541 9.69102 10.5426 9.48593 10.8793 9.3349C11.2161 9.18401 11.5857 9.10849 11.9882 9.10849C12.3909 9.10849 12.7626 9.18589 13.103 9.34085C13.2296 9.39835 13.3495 9.46321 13.4626 9.5356ZM13.1205 14.6652C12.7837 14.8163 12.4141 14.8917 12.0116 14.8917C11.6089 14.8917 11.2373 14.8143 10.8967 14.6595C10.7702 14.6018 10.6504 14.5369 10.5374 14.4644L14.4135 10.5883C14.4634 10.6852 14.5086 10.786 14.549 10.891C14.6843 11.2432 14.7521 11.6169 14.7521 12.0117C14.7521 12.4066 14.6843 12.7801 14.549 13.1325C14.4135 13.4847 14.2258 13.7904 13.9857 14.0497C13.7457 14.309 13.4574 14.5142 13.1205 14.6652ZM13.707 15.8671C14.2258 15.6426 14.6748 15.3388 15.0541 14.9554C15.4334 14.5723 15.7275 14.1252 15.9366 13.6143C16.1457 13.1033 16.2502 12.5614 16.2502 11.9886C16.2502 11.4156 16.1457 10.8736 15.9366 10.3626C15.8126 10.0596 15.66 9.77971 15.4789 9.52291L18 7.0018L16.9982 6L14.4694 8.52879C14.2424 8.37399 13.9959 8.24019 13.7302 8.12722C13.2114 7.90662 12.6386 7.79616 12.0116 7.79616C11.3844 7.79616 10.8116 7.9085 10.2928 8.13302C9.77409 8.35754 9.32504 8.66134 8.94572 9.04457C8.5664 9.42781 8.27215 9.87498 8.06315 10.3859C7.85415 10.8968 7.74963 11.4388 7.74963 12.0117C7.74963 12.5846 7.85415 13.1263 8.06315 13.6375C8.18724 13.9406 8.33983 14.2204 8.52097 14.4774L6 16.9982L7.0018 18L9.53043 15.4713C9.75746 15.6262 10.0039 15.76 10.2696 15.8729C10.7884 16.0936 11.3612 16.2038 11.9882 16.2038C12.6154 16.2038 13.1882 16.0915 13.707 15.8671Z"
				fill="white"
			/>
		</svg>
	),

	usdcRound: (props: LucideProps) => (
		<svg
			width="40"
			height="40"
			viewBox="0 0 40 40"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<circle cx="20" cy="20" r="20" fill="#2775CA" />
			<path
				fillRule="evenodd"
				clipRule="evenodd"
				d="M20 10C25.5229 10 30 14.4771 30 20C30 25.5229 25.5229 30 20 30C14.4771 30 10 25.5229 10 20C10 14.4771 14.4771 10 20 10ZM22.2849 12.9586C22.0594 12.8865 21.875 13.0204 21.875 13.2571V13.8394C21.875 13.9981 21.9947 14.1788 22.1437 14.2334C24.5372 15.1101 26.25 17.4107 26.25 20.1042C26.25 22.7976 24.5372 25.0982 22.1437 25.9749C21.9802 26.0348 21.875 26.1949 21.875 26.369V26.9512C21.875 27.1879 22.0594 27.3218 22.2849 27.2498C25.3097 26.2833 27.5 23.4496 27.5 20.1042C27.5 16.7588 25.3097 13.925 22.2849 12.9586ZM18.125 13.2571C18.125 13.0204 17.9406 12.8865 17.7151 12.9586C14.6903 13.925 12.5 16.7588 12.5 20.1042C12.5 23.4496 14.6903 26.2833 17.7151 27.2498C17.9406 27.3218 18.125 27.1879 18.125 26.9512V26.369C18.125 26.2102 18.0054 26.0295 17.8563 25.9749C15.4628 25.0982 13.75 22.7976 13.75 20.1042C13.75 17.4107 15.4628 15.1101 17.8563 14.2334C18.0054 14.1788 18.125 13.9981 18.125 13.8394V13.2571ZM20.3125 15.1042H19.6875C19.5149 15.1042 19.375 15.2441 19.375 15.4167V16.3848C18.136 16.561 17.335 17.3855 17.335 18.4433C17.335 19.8125 18.1627 20.3438 19.9098 20.5789C21.0951 20.7731 21.4325 21.0284 21.4325 21.7029C21.4325 22.3772 20.86 22.8268 20.0531 22.8268C18.9619 22.8268 18.6029 22.3494 18.4704 21.7314C18.4391 21.585 18.3137 21.478 18.164 21.478H17.4511C17.2715 21.478 17.1311 21.6397 17.162 21.8166C17.3435 22.8568 18.0116 23.62 19.375 23.805V24.7917C19.375 24.9643 19.5149 25.1042 19.6875 25.1042H20.3125C20.4851 25.1042 20.625 24.9643 20.625 24.7917V23.8045C21.9157 23.5995 22.7404 22.7009 22.7404 21.5904C22.7404 20.1293 21.8513 19.6388 20.1346 19.4039C18.8676 19.2199 18.6226 18.9235 18.6226 18.3309C18.6226 17.769 19.0516 17.3704 19.8793 17.3704C20.628 17.3704 21.0585 17.6316 21.2364 18.2342C21.2761 18.3686 21.3964 18.4637 21.5366 18.4637H22.194C22.3773 18.4637 22.52 18.2952 22.4809 18.1161C22.2728 17.1639 21.6311 16.5925 20.625 16.4123V15.4167C20.625 15.2441 20.4851 15.1042 20.3125 15.1042Z"
				fill="white"
			/>
		</svg>
	),
	cosmosPale: (props: LucideProps) => (
		<svg
			width="16"
			height="16"
			viewBox="0 0 16 16"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<g id="logo/cosmos">
				<path
					id="Vector"
					fillRule="evenodd"
					clipRule="evenodd"
					d="M9.46261 5.5356L5.58638 9.41198C5.53639 9.31499 5.49129 9.21393 5.45085 9.10914C5.31548 8.75691 5.2478 8.38323 5.2478 7.98857C5.2478 7.59358 5.31548 7.22003 5.45085 6.86784C5.58638 6.51548 5.77409 6.20979 6.01411 5.95031C6.25412 5.69102 6.54258 5.48593 6.87928 5.3349C7.21615 5.18401 7.58575 5.10849 7.98824 5.10849C8.39091 5.10849 8.76256 5.18589 9.10303 5.34085C9.2296 5.39835 9.34946 5.46321 9.46261 5.5356ZM9.12055 10.6652C8.78372 10.8163 8.41408 10.8917 8.01159 10.8917C7.60892 10.8917 7.23731 10.8143 6.89668 10.6595C6.77024 10.6018 6.65038 10.5369 6.53743 10.4644L10.4135 6.58834C10.4634 6.68516 10.5086 6.78605 10.549 6.89101C10.6843 7.24324 10.7521 7.61691 10.7521 8.01174C10.7521 8.40657 10.6843 8.78012 10.549 9.13247C10.4135 9.48466 10.2258 9.79035 9.98572 10.0497C9.74571 10.309 9.45741 10.5142 9.12055 10.6652ZM9.70703 11.8671C10.2258 11.6426 10.6748 11.3388 11.0541 10.9554C11.4334 10.5723 11.7275 10.1252 11.9366 9.61426C12.1457 9.10335 12.2502 8.56139 12.2502 7.98857C12.2502 7.4156 12.1457 6.87363 11.9366 6.36255C11.8126 6.05955 11.66 5.77971 11.4789 5.52291L14 3.0018L12.9982 2L10.4694 4.52879C10.2424 4.37399 9.99591 4.24019 9.7302 4.12722C9.21145 3.90662 8.63863 3.79616 8.01159 3.79616C7.38442 3.79616 6.8116 3.9085 6.29285 4.13302C5.77409 4.35754 5.32504 4.66134 4.94572 5.04457C4.5664 5.42781 4.27215 5.87498 4.06315 6.38589C3.85415 6.89684 3.74963 7.43876 3.74963 8.01174C3.74963 8.58455 3.85415 9.12635 4.06315 9.63747C4.18724 9.94064 4.33983 10.2204 4.52097 10.4774L2 12.9982L3.0018 14L5.53043 11.4713C5.75746 11.6262 6.00393 11.76 6.26963 11.8729C6.78843 12.0936 7.36124 12.2038 7.98824 12.2038C8.61542 12.2038 9.18823 12.0915 9.70703 11.8671Z"
					fill="#E5EEFF"
					fillOpacity="0.6"
				/>
			</g>
		</svg>
	),
	linkChainlink: (props: LucideProps) => (
		<svg
			width="40"
			height="40"
			viewBox="0 0 40 40"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<circle cx="20" cy="20" r="20" fill="white" />
			<g clip-path="url(#clip0_3858_27515)">
				<mask
					id="mask0_3858_27515"
					style={{ maskType: "luminance" }}
					maskUnits="userSpaceOnUse"
					x="8"
					y="8"
					width="24"
					height="24"
				>
					<path
						d="M20.0091 14.2376L25.0594 17.1085V22.8732L20.023 25.7625L14.9727 22.8962V17.1314L20.0091 14.2376ZM20.0091 10L18.1558 11.064L13.1101 13.9578L11.2568 15.0218V17.1406V22.9007V25.0195L13.1101 26.0743L18.1604 28.9452L20.0137 30L21.8671 28.936L26.9035 26.0422L28.7568 24.9828V22.8641V17.0993V14.9805L26.9035 13.9257L21.8532 11.0548L19.9999 10H20.0091Z"
						fill="white"
					/>
				</mask>
				<g mask="url(#mask0_3858_27515)">
					<rect x="8" y="8" width="24" height="24" fill="#2A5AD9" />
				</g>
			</g>
			<defs>
				<clipPath id="clip0_3858_27515">
					<rect
						width="24"
						height="24"
						fill="white"
						transform="translate(8 8)"
					/>
				</clipPath>
			</defs>
		</svg>
	),
	noActions: (props: LucideProps) => (
		<svg
			width="48"
			height="48"
			viewBox="0 0 48 48"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			{...props}
		>
			<path
				fillRule="evenodd"
				clipRule="evenodd"
				d="M28.5581 12.3081C28.8021 12.064 29.1979 12.064 29.4419 12.3081L34.4419 17.3081C34.686 17.5521 34.686 17.9479 34.4419 18.1919L29.4419 23.1919C29.1979 23.436 28.8021 23.436 28.5581 23.1919C28.314 22.9479 28.314 22.5521 28.5581 22.3081L32.4911 18.375H14C13.6548 18.375 13.375 18.0952 13.375 17.75C13.375 17.4048 13.6548 17.125 14 17.125H32.4911L28.5581 13.1919C28.314 12.9479 28.314 12.5521 28.5581 12.3081ZM19.4419 24.8081C19.686 25.0521 19.686 25.4479 19.4419 25.6919L15.5089 29.625H34C34.3452 29.625 34.625 29.9048 34.625 30.25C34.625 30.5952 34.3452 30.875 34 30.875H15.5089L19.4419 34.8081C19.686 35.0521 19.686 35.4479 19.4419 35.6919C19.1979 35.936 18.8021 35.936 18.5581 35.6919L13.5581 30.6919C13.314 30.4479 13.314 30.0521 13.5581 29.8081L18.5581 24.8081C18.8021 24.564 19.1979 24.564 19.4419 24.8081Z"
				fill="white"
			/>
		</svg>
	),
};
