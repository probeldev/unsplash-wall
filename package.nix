{
  buildGoModule,
  swaybg
}:
buildGoModule {
  name = "unsplashwall";
  src = ./.;
  vendorHash = null;

  buildInputs = [ swaybg ];
}
