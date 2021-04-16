package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.Gpu;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface GpuRepository extends JpaRepository<Gpu,Long> {
    List<Gpu> findAll();
}
