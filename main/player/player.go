components {
  id: "player"
  component: "/main/player/scrPlayer.script"
  properties {
    id: "Speed"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"ship\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/atlas.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "factory"
  type: "factory"
  data: "prototype: \"/main/shots/objBasicShot.go\"\n"
  ""
}
