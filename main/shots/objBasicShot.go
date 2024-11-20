components {
  id: "scrBasicShot"
  component: "/main/shots/scrBasicShot.script"
  properties {
    id: "speed"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"basic_shot\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/atlas.atlas\"\n"
  "}\n"
  ""
}
